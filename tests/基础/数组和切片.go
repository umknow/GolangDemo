package main

import (
	"fmt"
	"strconv"
)

func printSlince(s []int) {
	fmt.Printf("len=%d,cap=%d,v>>%v\n", len(s), cap(s), s)
}

func main() {

	// 切片
	s := make([]int, 2, 6)
	s = append(s, 2, 3, 4, 5,3, 3, 4, 5,3, 3, 4, 5,3, 3, 4,)
	//s := []int{2,3,4,5,7,12}
	printSlince(s)
	printSlince(s[:0])
	printSlince(s[:4])
	printSlince(s[2:])

	var s1 = []int{}
	s1 = append(s1, 1, 2)
	fmt.Println(s1)

	// 类型转换
	str1 := 1232
	fmt.Printf("【类型转换】str T >>%T,v>> %v\n", string(str1), string(str1))
	fmt.Printf("【类型转换】str T >>%T,v>> %v\n", strconv.Itoa(str1), strconv.Itoa(str1))


	// 测试引用类型改变原始数值测试
	var s2 = []int{1, 2, 3}
	s3 := s2
	s3[1] = 33
	fmt.Println("s2 >>", s2)
	fmt.Println("s3 >>", s3)

	// copy示例
	// 创建一个新的更大的切片并把原分片的内容都拷贝过来
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)
	n := copy(sl_to, sl_from)
	fmt.Println("【copy】num:", n, "items", sl_to)




	// 一维数组
	var ar = [...]string{"ss", "ttt", "ee"}
	fmt.Println(ar)
	var a2 = ar[1:]
	fmt.Printf("ar T >>%T, v >>%v\n", ar, ar)
	fmt.Printf("a2 T >>%T\n", a2)

	// 二维数组
	var arr = [3][4]int{{1, 2, 3, 4}, {3, 4, 5, 5}, {3, 5, 7, 4}}
	//b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Println(arr[i][j])
		}
	}

	// 二维切片
	var slince [][]interface{}
	s5 := []interface{}{
		1,2,3,"ss",
	}

	slince = append(slince, s5)
	fmt.Println(slince)

	}
