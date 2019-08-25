package main

import (
	"bufio"
	"github.com/dollarkillerx/easyutils/clog"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Reader interface {
	Read(chan []byte)
}

type Write interface {
	Write(chan []byte)
}

type LogProcess struct {
	rc    chan []byte
	wc    chan []byte
	read  Reader
	write Write
}

func (l *LogProcess) Process() {
	// 解析模块

	for {
		select {
		case data := <-l.rc:
			upper := strings.ToUpper(string(data))

			l.wc <- []byte(upper)
		}
	}
}

// 读取模块
type ReadFromFile struct {
	path string
}

func (r *ReadFromFile) Read(rc chan []byte) {
	// 打开文件
	file, e := os.Open(r.path)
	if e != nil {
		clog.Println(e.Error())
		panic(e.Error())
	}

	// 从文件末尾开始逐行读取内容
	file.Seek(0, 2)

	reader := bufio.NewReader(file)

	for {
		// 按照行读取
		bytes, e := reader.ReadBytes('\n')
		if e == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if e != nil {
			clog.Println(e.Error())
			panic(e.Error())
		}
		rc <- bytes
	}

}

// 解析模块
type WriteFromInfluxDb struct {
	influxDBDsn string
}

func (w *WriteFromInfluxDb) Write(wc chan []byte) {
	for {
		select {
		case data := <-wc:
			log.Println(data)
		}
	}
}

func main() {
	read := &ReadFromFile{
		path: "./tmp/access.log",
	}

	write := &WriteFromInfluxDb{
		influxDBDsn: "username@password",
	}

	lp := &LogProcess{
		rc:    make(chan []byte),
		wc:    make(chan []byte),
		read:  read,
		write: write,
	}

	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	ints := make(chan int)

	<-ints
}
