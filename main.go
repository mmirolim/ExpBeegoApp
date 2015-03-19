package main

import (
	_ "expbeego/routers"

	"github.com/astaxie/beego"
)

func main() {
	// init logger
	beego.SetLogger("file", `{"filename": "logs/test.log"}`)
	beego.Run()
}
