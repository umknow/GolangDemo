package main

import "fmt"

func main(){
	a := 100
	var ip *int
	fmt.Println(ip)

	// 用p接收指针类型
	var p = &a
	fmt.Printf("【a】T >> %T, V >> %v\n", a, a)
	fmt.Printf("【&a】T >> %T, V >> %v\n", &a, &a)
	fmt.Printf("【p】T >> %T, V >> %v\n", p, p)
	fmt.Printf("【&p】T >> %T, V >> %v\n", &p, &p)  // 指针的指针
	fmt.Printf("【*&p】T >> %T, V >> %v\n", *&p, *&p)  // 指针的数值
	fmt.Printf("【*p】T >> %T, V >> %v\n", *p, *p)
	fmt.Printf("【&*p】T >> %T, V >> %v\n", &*p, &*p)  // == &a == p

	fmt.Println("**********************************")
	const COUNT  = 3
	// 数组指针
	//var arr1 [3]int = [3]int{1, 2, 3}
	var arr1 *[COUNT]string = &[COUNT]string{"wwww", "ssss", "rrrrr"}
	fmt.Printf("【arr1】T >> %T, v >> %v\n", arr1, arr1)

	// 指针数组
	var arr2 [COUNT]*string
	for i:=0;i<COUNT;i++{
		arr2[i] = &arr1[i]
	}
	fmt.Printf("【arr2】T >> %T, v >> %v\n", arr2, arr2)
	fmt.Printf("【arr2】T >> %T, v >> %v\n", arr2, *arr2[1])


}
