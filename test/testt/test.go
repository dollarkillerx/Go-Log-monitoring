/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-10
* Time: 下午7:36
* */
package main

import "fmt"

type ren struct {
	car cats
}

type cats interface {
	Run()
}

type car struct {
	CatName string
}

func (c *car) Run()  {
	fmt.Println("Run")
}

func main() {
	c := &car{CatName:"本田黑科技"}
	r := ren{car: c}
	fmt.Println(r)
}
