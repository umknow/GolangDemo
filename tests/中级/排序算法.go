package main

import "fmt"

func main() {
	var arr = [7]int{1,6, 5, 4, 2, 7, 9}
	fmt.Println("arr len >>", len(arr))
	bulle_sort(arr) // 冒泡排序
}

func bulle_sort(lis [7]int) { // 冒泡排序
	var lis_len = len(lis) - 1
	for i := 0; i < lis_len; i++ {
		flag := true  // 【优化】减少遍历次数
		for j := 0; j < lis_len-i; j++ {
			if lis[j] > lis[j+1] {
				lis[j], lis[j+1] = lis[j+1], lis[j]
				flag = false
				}
		}
		if flag{
			break
		}
	}
	fmt.Println("【冒泡排序结果】", lis)
}
