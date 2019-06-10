/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-10
* Time: 下午6:03
* */
package main

import (
	"fmt"
	"strings"
)

type Reader interface {
	Read(chan string)
}

type Write interface {
	Write(chan string)
}

type LogProcess struct {
	rc chan string
	wc chan string
	read Reader
	write Write
}

type ReadFromFile struct {
	path string
}

func (r *ReadFromFile)Read(rc chan string)  {
	msg := "message"
	rc <- msg
}

func (l *LogProcess)Process()  {
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}

type WriteToInfluxDb struct {

}

func (w *WriteToInfluxDb)Write(wr chan string)  {
	data := <-wr
	fmt.Println(data)
}


func main() {
	r := &ReadFromFile{
		path:"/tmp/access.log",
	}

	w := &WriteToInfluxDb{}

	process := &LogProcess{
		rc:    make(chan string),
		wc:    make(chan string),
		write: w,
		read:  r,
	}
	process = process
}
