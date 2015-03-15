package controllers

import "errors"

type MainController struct {
	BaseCtrl
}

func (c *MainController) Get() {
	err := errors.New("internal error")
	c.onErrReturnJson(err, internalErrorJson)
}

func (c *MainController) Post() {
	err := errors.New("some error")
	c.onErrReturnJson(err, dbErrorJson)
	//	c.abortOnErr(err, "dbError")
}
