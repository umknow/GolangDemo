/*
	指针的使用
*/
package main

import "fmt"

func demo(a,b *int)int  {
	*a = 5
	*b = 6
	return  *a * *b //这里出现连续的两个*，Go编译器会根据上下文自动识别乘法与两个引用
}


func main()  {
	a := 3
	b := 4
	p1 := &a
	p2 := &b
	fmt.Printf("%v, %v", *p1, *p2)
	fmt.Println(demo(p1, p2))

}