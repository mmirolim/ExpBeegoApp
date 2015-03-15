package main

import (
	"encoding/json"
	"expbeego/controllers"
	_ "expbeego/routers"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
)

func main() {
	beego.Errorhandler("dbError", dbError)
	beego.Run()
}

// custom error handler when Aborting request
// problem with setting Headers.
func dbError(rw http.ResponseWriter, r *http.Request) {
	rj := controllers.RJson{Msg: "DB_ERROR_INTERNAL"}
	data, err := json.Marshal(&rj)
	if err != nil {
		data = []byte(`{msg:"json_marshal_error"}`)
	}
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(rw, string(data))
}
