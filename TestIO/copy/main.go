/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-10
* Time: 下午9:20
* */
package main

import (
	"bufio"
	"context"
	"io"
	"os"
	"sync"
)

func main() {
	copy()
}


func copy() {
	old, e := os.OpenFile("./test.mp4", os.O_RDONLY, 00666)
	new, e := os.OpenFile("./new.mp4", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 00666)
	if e != nil {
		panic(e.Error())
	}

	reader := bufio.NewReader(old)
	writer := bufio.NewWriter(new)
	bytes := make([]byte, 1024)

	for {
		_, e := reader.Read(bytes)
		if e == io.EOF {
			break
		}else{
			writer.Write(bytes)
		}
	}
	defer func() {
		writer.Flush()
		old.Close()
		new.Close()
	}()
}

func read(ch chan []byte,cancelFunc context.CancelFunc,wg *sync.WaitGroup)  {
	old, e := os.OpenFile("./test.mp4", os.O_RDONLY, 00666)
	new, e := os.OpenFile("./new.mp4", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 00666)

	defer func() {
		old.Close()
		new.Close()
	}()
	if e != nil {
		panic(e.Error())
	}

	reader := bufio.NewReader(old)
	writer := bufio.NewWriter(new)
	bytes := make([]byte, 1024)

	for {
		_, e := reader.Read(bytes)
		writer.Write(bytes)

		if e == io.EOF {
			cancelFunc()
			break
		}
	}
	wg.Done()

}

func write(ctx context.Context,ch chan []byte,wg *sync.WaitGroup)  {
	file, e := os.OpenFile("./new.mp4", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 00666)
	defer func() {
		file.Close()
	}()
	if e != nil {
		panic(e.Error())
	}

	writer := bufio.NewWriter(file)
	forloop:
	for  {
		select {
		case <-ch:
			writer.Write(<-ch)
		case <-ctx.Done():
			break forloop
		}
	}
	writer.Flush()
	wg.Done()

}
