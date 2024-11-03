package main

import "fmt"

func main55() {
	var array  = [1] int  {1}

	var array1  = [2] int {1,2}

	var array2  = [3] int {1,2,3}
	
	var array3  = [] int {1,2,3,4}

	var array4  = [5] int {1,2,3,4,5}

	fmt.Println(array)

	fmt.Println(array1)

	fmt.Println(array2)

	fmt.Println(array3)
	fmt.Println(array4)
/*	for true {
		fmt.Println("opps")
	}*/
	i := 0
	for i = 0 ; i <= 10 ; i++ {
		fmt.Println("this is your number ", i)
		if i == 6 {
			continue
		}
	

	}



}
func opps(num int, num1 int) int {
	return num*num1
}
func oppse() {
	var a, b = 4 ,5
	var (
		s = 4
		d = 5
	)
	fmt.Println(a,b,s,d)
	const red , cb = 789 , "kasmir"
	//---------------------------------
	const (
		der = 78
		bc = 79
	)
	//---------------------------------
	if der == bc {
		fmt.Println("yes")

	}else if der >= bc{
	fmt.Println("yes")
	}else {
		fmt.Println("noooooo")
	} 
	var h2r = 0
	switch h2r{
	case 0:
		fmt.Println("yse")
	case 1:
		fmt.Println("no")
	case 2:
		fmt.Println("no")
	case 3:
		fmt.Println("no")
	case 4:
		fmt.Println("no")
	case 5:
		fmt.Println("no")
	default:
		fmt.Println("not found")
	}


		
}

