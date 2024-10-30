package main

import "fmt"

func main() {
	fmt.Println("")
	// dba n3afo b Array in go
	// Array declaration
	var num int = 8
	for i := 0; i < num; i++ {
		var Array_1 = [6] int{8, 2, 6, 4, 0, 6}
		var Array_2 = [6] int{7, 2, 1, 9, 9, 6}
		var Array_3 = [6] int{6, 2, 1, 2, 5, 6}
		// had dari9a bach declarai b array
		fmt.Println("this is array", Array_1)
		fmt.Println("this is array", Array_2)
		fmt.Println("this is array", Array_3)
		fmt.Println("------------------------")
		// array access specific element
		var arr = [3] string {"ahmed", "mohamed", "ali"}
		fmt.Println("the first name in the array is", arr[1])
		// change Array element
		var arr1 = [3] string {"ahmed", "mohamed", "aaa1li"}
		arr1[2] = "harmsh"
		fmt.Println("the first name in the array is", arr1)
		arrry()
	}
}
/*func Array(harmsh int) {
	fmt.Println("hashlmao")
	harmsh = 0
	for i := 10; i < harmsh; harmsh++ {
		fmt.Println(harmsh)

}*/
func arrry() {
	fmt.Println("------------------------")
	var num =[4] int {1, 2, 3, 5}
	fmt.Println("this is arrry", num)
}