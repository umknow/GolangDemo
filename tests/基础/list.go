/*
	list的测试
	本质：链表
*/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	optList1()
	optList2()
}

func optList1()  {
	var list1 list.List
	// 填充数据
	list1.PushFront("xyz")  // 在头部添加
	list1.PushBack("baks")  // 在尾部添加
	list1.PushBack([]int{1,2})  // 在尾部添加不同数据类型
	printListInfo("【list1】声明", list1)
	iterateList("【list1】", list1)

	// 清除
	list1.Init()
	printListInfo("【list1】清除后", list1)

}

func optList2()  {
	list2 := list.New()  // 已经过初始化相当于map make()后 ；var 声明的全为nil
	list2.PushFront("ssdsffd")

	element := list2.PushBack("sssss")
	list2.InsertAfter("wwww", element)

	printListInfo("【list2】声明", *list2)
	iterateList("【list2】", *list2)
	iterateListReverse("【list2】", *list2)

	list3 := new(list.List)
	printListInfo("【list3】声明", *list3)

	// 测试list是值类型还是引用类型（本身是值类型，new（）后返回的是指针类型）
	list4 := *list2
	list4.PushBack("wangyongsheng")
	printListInfo("【list2_】声明", *list2)
	printListInfo("【list4_】声明", list4)
	}


func printListInfo(info string, l list.List) {
	fmt.Printf("%v :%T, %v\t,长度：%d\n", info,l, l, l.Len())
}

func iterateList(info string, l list.List) {  // 顺序遍历
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v iter: %v \n", info, e.Value)
	}
}

func iterateListReverse(info string, l list.List) {  // 倒序遍历
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Printf("%v iter: %v \n", info, e.Value)
	}
}
