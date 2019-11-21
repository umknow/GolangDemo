package main

import "fmt"

func main()  {
	str := "时间asdf我的"
	fmt.Println([]byte(str))
	fmt.Println([]rune(str))
	fmt.Println(string(str[3]))
	fmt.Println("长度>>", len(str))
	//遍历
	for i, ch:= range str{
		fmt.Printf("%d, %c\n", i, ch)
	}
	for i,ch := range []byte(str){
		fmt.Printf("%d, %c\n", i, ch)
	}
	for i,ch := range []rune(str){
		fmt.Printf("%d, %c\n", i, ch)
	}

	// 相关操作


}

