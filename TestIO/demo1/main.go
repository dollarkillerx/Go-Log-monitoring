/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-10
* Time: 下午8:59
* */
package main


import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	write()
}

func read()  {
	// 获得文件句柄 (O_RDONLY) 只读
	file, e := os.OpenFile("./test.file", os.O_RDONLY, 0666)
	defer file.Close()
	if e != nil {
		panic(e.Error())
	}

	// 创建读取器
	reader := bufio.NewReader(file)

	// 创建缓冲区
	bytes := make([]byte, 1024) // 1024为1K

	for {
		_, e := reader.Read(bytes)
		fmt.Printf("%s",bytes)
		if e == io.EOF{
			break
		}
	}
}

func write()  {
	// 获取文件句柄   (O_CREATE) 不存在就创建 (O_WRONLY) 写入  (O_APPEND) 追加
	file, e := os.OpenFile("./new.text", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 00666)
	defer file.Close()
	if e != nil {
		panic(e.Error())
	}

	// 创建写出器
	writer := bufio.NewWriter(file)

	_, e = writer.Write([]byte("hello"))
	if e != nil {
		fmt.Printf(e.Error())
	}

	writer.Flush() //倒干缓冲区
}