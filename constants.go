package main

import "fmt"

func main1() {
	// loyma gha n3afo b "const" ou ly haya 9im tabita y3ni aks nta3 variable
	// how to declare a single constant 
	const name string = "ramy"
	fmt.Println(name)
	// how to declare multiple constants in const wah7da
	const (
		name1 string = "ramy"
        name2 int = 544
        name3 bool = true
	)
	fmt.Println(name1, name2, name3)
}