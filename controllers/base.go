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

// always serve in json format
// does not work for custom aborts
func (c *BaseCtrl) Finish() {
	c.ServeJson()
}

// this is beego feature added in 1.4.3
// custom error handler but you can't set headers
// (or I could not do it)
func (c *BaseCtrl) abortOnErr(err error, errHandler string) {
	if err != nil {
		beego.Warn(err)
		c.Abort(errHandler)
	}
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
	w := c.Ctx.ResponseWriter
	// set content-type to json
	// if needed set other headers here
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(cerr))
}
