/*
	递归测试
	注意事项：注意栈的溢出
	内存主要使用的栈， 堆
*/
package main

import "fmt"


func factoria(n int) int{
	if n== 0{
		return 1
	}
	return n*factoria(n-1)
}


func main()  {
	fmt.Println(factoria(6))
}

