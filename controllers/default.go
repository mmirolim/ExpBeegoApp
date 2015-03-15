package controllers

import (
	"errors"

	"github.com/astaxie/beego"
)

type MainController struct {
	BaseCtrl
}

func (c *MainController) Get() {
	err := errors.New("internal error")
	// return internal error in json format
	c.onErrReturnJson(err, internalErrJson)
}

func (c *MainController) Post() {
	err := errors.New("some error")
	// return db error response in json format
	c.onErrReturnJson(err, dbErrJson)
	beego.Error("this will not print")
}
