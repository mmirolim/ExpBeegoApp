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

type CtrlError string

var (
	dbErrorJson       CtrlError
	internalErrorJson CtrlError
)

func init() {
	// parse json err response
	jsonParseError := []byte(`{msg:"JSON_PARSE_ERROR"}`)
	data, err := json.Marshal(RJson{Msg: "DB_ERROR"})
	if err != nil {
		data = jsonParseError
	}
	dbErrorJson = CtrlError(data)
	data, err = json.Marshal(RJson{Msg: "INTERNAL_ERROR"})
	if err != nil {
		data = jsonParseError
	}
	internalErrorJson = CtrlError(data)
}

func (c *BaseCtrl) Prepare() {
	c.Data["title"] = "EXP.App"
}

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
	if err == nil {
		return
	}
	// log error
	beego.Error(c.Ctx.Request.URL, err)
	w := c.Ctx.ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(cerr))
}
