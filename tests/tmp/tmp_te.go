package main


import "fmt"

type Ob struct{
	name string
}

func main()  {
	fmt.Printf("%T\n", Ob{})

	d := make(map[interface{}]interface{})
	d[Ob{}] = "sdsf"
	d[12] = "f"
	fmt.Println(d)
}

