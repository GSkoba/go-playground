package main

import (
	"fmt"
	"time"
)

func main()  {
	i := 2 
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3: 
		fmt.Println("tree")
	}

	fmt.Println(time.Now().Date())

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: 
		fmt.Println("Its the weekend")
	default: 
		fmt.Println("It a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It before noon")
	default:
		fmt.Println("Its after noon")
	}

	whatAmI := func(i interface{})  {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am bool")
		case int:
			fmt.Println("I am int")
		default:
			fmt.Print("Dont know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}