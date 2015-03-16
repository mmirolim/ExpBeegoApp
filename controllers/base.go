package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

type RJson struct {
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}

type BaseCtrl struct {
	beego.Controller
}

// define error in Ctrls
type CtrlError string

var (
	// ctrlerrors should be predeclared
	// so will be able to change them from one place
	dbErrJson       CtrlError
	internalErrJson CtrlError
)

func init() {
	// prepare marshaled response for db_error
	data, err := json.Marshal(RJson{Msg: "DB_ERROR"})
	if err != nil {
		data = []byte(fmt.Sprintf("%s", err))
	}
	dbErrJson = CtrlError(data)

	// prepare marshaled response for internal_error
	data, err = json.Marshal(RJson{Msg: "INTERNAL_ERROR"})
	if err != nil {
		data = []byte(fmt.Sprintf("%s", err))
	}
	internalErrJson = CtrlError(data)
}

func (c *BaseCtrl) Prepare() {
	c.Data["title"] = "EXP.App"
}

// stop execution of response if there is error
// log it and respond with predefined message in json format
func (c *BaseCtrl) onErrReturnJson(err error, cerr CtrlError) {
	// early return if no error
	if err == nil {
		return
	}
	// don't forget loggin errors
	// to get more data we can extract it from Context
	// here we loggin errors with URL so we will know where it happened
	beego.Error(c.Ctx.Request.URL, err)
	// set content-type to json
	// if needed set other headers here
	w := c.Ctx.ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "MyAPP")
	// set proper error status
	// and it should be set before
	// call to ResponseWriter.Write
	// otherwise status will be set to htt.StatusOk
	// http://godoc.org/net/http#ResponseWriter
	switch cerr {
	case dbErrJson:
		w.WriteHeader(400)
	default:
		w.WriteHeader(500)
	}
	// write response
	w.Write([]byte(cerr))
	// maybe not to panic to get execution time
	c.StopRun()
}

// create struct here just
// to make it easier see assocsiated struct
// Mail Notifier config
type NotifyMailConfig struct {
	Apikey    string        //Mandrill Api Key
	WaitTime  time.Duration // Wait time if internet connnection epsont wait for this time
	Timeout   time.Duration // If Mandrill Api dont responce too long close connection in this time
	BuffLimit int           // Buffer Limit
	Retry     int           //If error occures while sending retry it n times
	RatePerH  int           //Mandrill Send limit per Hour
}

type ConfParam struct {
	param  *int
	key    string
	errStr string
}

// init some service
func initService() {
	var NFS NotifyService
	//Init Notification Service Config
	var err error
	var wt, tm, buffLimit, retry, hrlimit, groupId int

	// conf to read
	confParams := []ConfParam{
		{&wt, "wait_time", "wt error"},
		{&tm, "timeout", "timeout error"},
		{&buffLimit, "buff_limit", "buf error"},
		{&retry, "retry", "retry error"},
		{&hrlimit, "hour_limit", "hl error"},
		{&groupId, "group_id", " group id error"},
	}
	// iterate conf
	// it could be further simplified
	for _, v := range confParams {
		*v.param, err = beego.AppConfig.Int(v.key)
		if err != nil {
			beego.Error(v.errStr, err)
			// stop iterating if there is error
			return
		}
	}

	conf := NotifyMailConfig{
		WaitTime:  time.Duration(wt) * time.Second,
		Timeout:   time.Duration(tm) * time.Second,
		BuffLimit: buffLimit,
		Retry:     retry,
		RatePerH:  hrlimit,
	}

	//Init Notification Service
	NFS.Start(conf)

}

type NotifyService struct{}

func (n *NotifyService) Start(conf NotifyMailConfig) {}
