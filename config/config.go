/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-7
* Time: 下午10:41
* */
package config

import (
	"Go-Log-monitoring/defs"
	"encoding/json"
	"os"
)

var(
	Config *defs.Config
)

func init() {
	file, e := os.Open("./config.json")
	if e != nil {
		panic(e.Error())
	}
	Config = &defs.Config{}
	decoder := json.NewDecoder(file)
	e = decoder.Decode(Config)
	if e != nil {
		panic(e.Error())
	}
}
