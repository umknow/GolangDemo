/*
	空接口的使用
	可定义实现不定长参数
*/
package main

import (
	. "fmt"
)

type A interface {

}

type Cat struct {
	name string
	age int
}

type Pers struct {
	name string
	sex string
}

func main() {
	// 赋值给空接口对象后，不能通过a1访问内部属性，需要对a1进行转型
	var a1 A = Cat{"加菲", 2}
	var a2 A = Pers{"Steven", "男"}
	var a3 A = "Learn golang !"
	var a4 A = 100
	var a5 A = 3.14
	showInfo(a1)
	showInfo(a2)
	showInfo(a3)
	showInfo(a4)
	showInfo(a5)
	Println("--------------------------")

	// 定义map， value是任何数据类型
	map1 := make(map[string]interface{})
	map1["name"] = "wang"
	map1["age"] = 13
	map1["height"] = 1.73
	Println(map1)
	Println("--------------------------")

	//定义切片，元素是任何数据类型
	s1 := make([]interface{}, 0)
	s1 = append(s1, 12, "sfsdf", []int{})
	Println(s1)

	// instance := a1
	transInterface(s1)
}


// 接口对象转型
// 接口对象.(type), 配合switch ...case
func transInterface(s []interface{})  {
	for i,_:=range s{
		Println("第", i+1, "个数据：")
		switch t:= s[i].(type) {
		case Cat:
			Printf("Cat对象，%s, %d\n",t.name, t.age)
		case Pers:
			Printf("Person对象，%s, %s\n",t.name, t.sex)
		// ...
			}
	}
}


func showInfo(a A)  {
	Printf("%T, %v\n", a, a)
}