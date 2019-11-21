/*
	判断练习
*/
package main

import (
	"fmt"
)

func score_grade(grade string)  {
	//grade := "A"
	switch grade {
	case "A":
		fmt.Println("grade is A!")
	case "B":
		fmt.Println("grade is B")
	default:  // else
		fmt.Println("grade is poor")
	}

	if grade == "A"{
		fmt.Println("[if]grade is A")
	}else{
		fmt.Println("not grade ...")
	}

}


func main()  {
	//score_grade("A")
	for i:=0;i<10;i++ {
		fmt.Println(i)
	}

	i := 0
	for ;;{
		if i >20{
			break
		}
		i++
		fmt.Printf("%d ",i)
	}


	for ;i<40;i++ {
		fmt.Println(i)
	}


	for{
		if i>30{
			fmt.Println("sssssssssss")
			break
		}
		i++
	}

	for i<50{
		fmt.Printf("%d ", i)
		i++
	}

	lis := [4]int{100, 200, 300}
	//fmt.Println(len(lis))
	//fmt.Printf("\n%T", lis)
	for j, value := range lis{
		fmt.Println("\n", j,":", value)
	}  // 这里遍历有点怪

	fmt.Println("____________________")
	str := "asdfasdfas"
	for i,v := range str{
		fmt.Printf("%d, %v, %c\n", i, v, v)
	}


	// 切片
	sli := []int{1, 2, 3}
	fmt.Printf("%d, ")
	}
