package main

import "fmt"

func main2() {
	fmt.Println("Hello World!")
	// dars nta3 lyoma how variables ou kifach ka yakhdmo
	//single variable declaration
	var name string = "ali"
	name = "yassine"
	fmt.Println(name)
	//multiple variables declaration in one line
	var name1, name2 string = "ali", "yassine"
	fmt.Println("your name is", name1, "and your name is", name2)
	//multiple variables declaration in one line without the same type
	var name3, name4 = "ali", 20
    fmt.Println("your name is", name3, "and your age is", name4)
    //tari9 bach declarai b ktar man variable fi **var** wha7da
	var(
		name5 = "varone"
		name6 = 666
		boolen = false
	)
	fmt.Println(name5,name6,boolen)
}
