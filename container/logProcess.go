/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-10
* Time: 下午4:40
* */
package container

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type LogProcess struct {
	Rc chan []byte // 读取chan
	Wc chan []byte // 写入chan
	Read Reader
	Write Writer
}

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan []byte)
}

func NewLogProcess(read Reader,write Writer) *LogProcess {
	return &LogProcess{
		Rc:make(chan []byte),
		Wc:make(chan []byte),
		Read:read,
		Write:write,
	}
}

// 读取模块
type ReadFromFile struct {
	Path string // 读取文件的路径
}

func (l *ReadFromFile) Read(rc chan []byte) {
	// 打开文件
	file, e := os.Open("./tmp/access.log")
	defer file.Close()
	if e != nil {
		panic(e.Error())
	}
	// 文件指针指向文件末尾 从末尾开始读
	file.Seek(0,2)

	// 从文件末尾开始卒行读取
	reader := bufio.NewReader(file)

	for  {
		bytes, e := reader.ReadBytes('\n')
		if e == io.EOF{
			time.Sleep(200 * time.Millisecond)
			continue
		}else if e != nil {
			log.Fatal(e.Error())
		}
		fmt.Printf("log: %s",bytes)
		rc <- bytes
	}
}

// 解析模块
func (l *LogProcess) Process() {
	for {
		select {
		case data := <- l.Rc:
			fmt.Printf("%s",data)
			l.Wc <- []byte(strings.ToUpper(string(data)))
		}
	}
}


// 写入模块
type WriteToInfluxDb struct {}

func (l *WriteToInfluxDb) Write(wc chan []byte) {
	for {
		select {
		case data := <- wc:
			fmt.Printf("%s",data)
		}
	}
}
