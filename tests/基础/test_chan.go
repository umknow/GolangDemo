/*
	Go 缓冲区的搭建
	目的：提高并发吞吐能力
	异步通信，提高生产者消费者效率
	并行
*/
package main

import (
	. "fmt"
	g "probep"
)

type Aiy struct {
	id int
	name string
}


func Producer(send chan <-int)  {  // 只写
	for i :=0;i<10;i++{
		send<-i  // 只写的方式
	Println("生产者 >>", i)
	}
	close(send)
}

func Consumer(recv <-chan int)  { // 只读
	for num:=range recv{
		Println("消费者 >>", num)
	}
}

func async(a Aiy){
	a.id++
	a.name = "wang"
	Println("async >>", a)
	Printf("async T >> %T\n", a)
}





func main()  {

	a := Aiy{1, "test"}
	//async(&a)  // 指针传递，改变原始数据
	async(a)
	Println("main >>", a)

	str := new(string)
	Println(str)

	// 创建缓冲区
	ch := make(chan int, 5)
	go Producer(ch)
	Consumer(ch)
	defer func(){  // 匿名函数的使用
		Println("defer func lambada")
	}()
	Println("closed ....")
	Println(g.Goperp())
	Println(g.Peoperq)


	}
