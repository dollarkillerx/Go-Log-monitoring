/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-7
* Time: 下午10:42
* */
package main

import (
	"Go-Log-monitoring/config"
	"Go-Log-monitoring/container"
	"time"
)

func main() {
	r := &container.ReadFromFile{
		Path: config.Config.Path,
	}
	w := &container.WriteToInfluxDb{}
	process := container.NewLogProcess(r, w)

	go process.Read.Read(process.Rc)
	go process.Process()
	go process.Write.Write(process.Wc)

	for {
		time.Sleep(time.Second*3)
	}
}
