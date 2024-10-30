package main

import "fmt"

func main3() {
	// danba gha n3afo b if condition 
	// ou else if condition 
	// ou else condition
	// ou else condition
	var (
		wow1 int = 100
		wow2 int = 10
	)
	if wow1 >= wow2 {
		fmt.Println("wow1 is greater than wow2")

	} else if wow1 == wow2 {
		fmt.Println("wow1 is equal to wow2")
	} else {
		fmt.Println("wow1 is less than wow2")
	}
	// ou dba 3anda Nested if condition ou chorot motadakhla
	var sleep string
	var sport string
	var study string
	var work string
	fmt.Println("do you sleep 8 hours a day? (yes/no)")
	fmt.Scanln(&sleep)
	if sleep == "yes" {
		fmt.Println("do you doing sports? (yes/no)")
		fmt.Scanln(&sport)
		if sport == "yes" {
			fmt.Println("do you study good? (yes/no)")
			fmt.Scanln(&study)
			if study == "yes" {
				fmt.Println("do you work 8 hours a day? (yes/no)")
				fmt.Scanln(&work)
				if work == "yes" {
					fmt.Println("your life and salary is vary good")
				} else {
					fmt.Println("thank you for your time")
				}
			}
		} else {
			fmt.Println("thank you for your time")
		}
	} else {
		fmt.Println("thank you for your time")
	}
}