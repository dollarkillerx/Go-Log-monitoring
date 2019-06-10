/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-7
* Time: 下午10:42
* */
package defs

type Config struct {
	Path string `json:"path"`
	InfluxDBDsn string `json:"influxDBDsn"`
	Host string `json:"host"`
}