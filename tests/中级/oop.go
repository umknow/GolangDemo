/*
	go的面向对象编程
*/
package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}


func (person *Person) showInfo() {
	fmt.Printf("My name is %s , age is %d ",person.name,person.age)
}

func (person *Person) setAge(age int) {
	person.age = age
	fmt.Println("my age is ", age)
}



type Student struct {
	Person  // 继承
	score int
	id int
}

func (student *Student) showInfo() {
	fmt.Println("I am a student ...")
}

func (student *Student) read() {
	fmt.Println("read book ...")
}




/*
	多态的demo
*/
type Human interface {
	speak(language string)
}

type Chinese struct {

}

type American struct {

}

func (ch Chinese) speak(language string ) {
	fmt.Printf("speck %s\n",language)
}

func (am American ) speak(language string ) {
	fmt.Printf("speck %s\n",language)
}




func main()  {
	// 继承
	student := Student{Person{"jake",16},1001,99}
	student.showInfo()
	student.setAge(22)
	student.read()
	fmt.Println("* *")

	// 多态
	var ch Human
	var am Human
	ch = Chinese{}
	am = American{}
	ch.speak("Chinese")
	am.speak("English")

	}
