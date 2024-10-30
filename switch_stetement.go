package main

import "fmt"

func main5() {
	fmt.Println("Hello World!")
	// dba gha n3afo b switch stetement 
	var week int 
	fmt.Scanln(&week)

	switch week {
	case 1:
		fmt.Println("sunday")
	case 2:
		fmt.Println("monday")
	case 3:
	    fmt.Println("tuesday")
	case 4:
		fmt.Println("wednesday")
	case 5:
		fmt.Println("thursday")
	case 6:
		fmt.Println("friday")
	case 7:
		fmt.Println("saturday")
	default:
		fmt.Println("your week not found")
	}
	// multiple case
	switch week {
    case 1, 7:
        fmt.Println("weekend")
    default:
        fmt.Println("week day")
    }
    // nested switch
    switch week {
    case 1, 7:
        fmt.Println("weekend")
    default:
        switch {
        case week >= 1 && week <= 5:
            fmt.Println("week day")
        default:
            fmt.Println("your week not found")
        }
    }
}