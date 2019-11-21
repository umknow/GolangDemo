package main

import (
	"fmt"
)


func addSum(nums ...int) (sum int) {  // 先定义一个sum变量然后再内部改变其值
	fmt.Printf("%T \n", nums)
	for _,value := range nums{
		sum += value
	}
	return // == return sum  !
	}

func main()  {
	slice:=[]int{1,2,3,4}
	add := addSum(slice...)  // 对切片拆包
	//add := addSum(1,2,3,4,5)  // 切片
	fmt.Println(add)

	s0 := new([3]int)
	for _,va := range *s0{
		fmt.Println(">>", va)
	}
	fmt.Printf(" s0 T>> %T", s0)


	a := [3]int{1, 2, 3}
	s1 := make([]int, 0, 3)
	for i := 0; i < cap(s1); i++ {
		s1 = append(s1, i)  // append数据到数组中
	}
	s2 := s1
	for i := 0; i < len(a); i++ {
		s1[i] = s1[i] + 1
	}

	fmt.Println(s1) //[1 2 3]
	fmt.Println(s2) //[1 2 3]
	}
