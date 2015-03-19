package controllers

import (
	"io/ioutil"
	"strings"
)

type Debug struct {
	BaseCtrl
}

func (c *Debug) Get() {
	// tpl for logs
	c.TplNames = "logs.tpl"
	// read file
	buf, err := ioutil.ReadFile("logs/test.log")
	c.onErrReturnJson(err, internalErrJson)
	str := string(buf)
	logs := strings.Split(str, "\n")
	n := len(logs)
	// slice for reversed logs
	reversedLogs := make([]string, n)
	for i := 0; i < n; i++ {
		reversedLogs[n-i-1] = logs[i]
	}

	c.Data["logs"] = reversedLogs
}
