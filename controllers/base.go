package controllers

import (
	"encoding/json"
	"fmt"

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
