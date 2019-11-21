/*
	闭包测试
*/
package main

import "fmt"

type nei func(num int) int

func adder() nei {
	sum := 0
	res := func(num int) int{
		sum += num
		return sum
	}
	fmt.Printf("adder res %T >>\n", res)
	return res
}


func main()  {
	res := adder()
	fmt.Printf("res >> %T\n", res)
	for i:=0;i<5;i++{
		fmt.Printf("i=%d \t", i)
		fmt.Println("res >>", res(i))
	}
}
