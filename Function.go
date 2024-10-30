package main

import "fmt"


func main6() {
	fmt.Println("Hello World!")
	// daba gha n3afo b function declaration
	// function declaration
	func() {
        fmt.Println("Hello World!")
    }() // function call
    // function declaration with parameters
    func(name string) {
        fmt.Println("Hello", name)
    }("World!")
    // function declaration with return value
    func() string {
        return "Hello World!"
    }() // function call
    // function declaration with parameters and return value
    func(name string) string {
        return "Hello " + name
    }("World!")
    // function declaration with parameters and return value
    func(name string) string {
        return "Hello " + name
    }("World!")
    // function declaration with parameters and return value
    func(name string) string {
        return "Hello " + name
    }("World!")
    // function declaration with parameters and return value
    func(name string) string {
        return "Hello " + name
    }("World!")


}