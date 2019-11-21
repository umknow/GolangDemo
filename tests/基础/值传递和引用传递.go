/*
	深度解析值传递和引用传递
*/

package main

import . "fmt"

func main()  {
	str := "wang"
	var p = &str

	Printf("p %v\n", *&p)
	Printf("*p %v\n", *p)
	//Printf("p %v\n", &p)

	action(&*&p)
	Printf("str new>>%T, %v\n", str, str)
	Printf("*p new >>%T, %v\n", p, *p)
}


func action(st **string)  {
	s := "ssssss"
	var p2 = &s
	Println("p2 >>", p2)
	*st = p2
	//*st = s
	Printf("st1>> %v, T>> %T\n", *&st, st)

}