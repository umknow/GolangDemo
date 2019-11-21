/*
	map的使用
*/
package main

import (
	"fmt"
)

func main()  {
	// 声明
	var dic0 = new(map[string]string)
	dic := make(map[string]string)
	fmt.Println(dic0, dic)
	fmt.Printf("%T, %T\n", dic0, dic)

	// 添加数据
	dic["email"] = "sssss"
	dic["name"] = "wys"
	dic["sss"] = "sssssssssss"

	// 删除元素
	delete(dic, "sss")
	fmt.Println(dic)

	// 查询&判断
	if value,ok:=dic["name"];ok{
		fmt.Println("已获取到值>>", value)
	}else {
		fmt.Println("未获取到指定值！", ok)
	}

	// 遍历
	iterateMap(dic)
	// 清空
	dic = *dic0
	fmt.Println("dic 清空后 >>", dic)
}

func iterateMap(d map[string]string)  {
	for i, value := range d{
		fmt.Println("【iter map】", i, value)
	}
}
