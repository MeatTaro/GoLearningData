package main

import (
	"fmt"
	"time"
)

func main()  {
	now := time.Now()
	weekOfDay := now.Weekday()
	fmt.Println(weekOfDay)
	switch weekOfDay {
	case 6, 7 : fmt.Println("Weekend")
	case 5: fmt.Println("Thank God. It's Friday")
	case 3: fmt.Println("Hump day")
	default: fmt.Println("Work Day")
		
	}
}