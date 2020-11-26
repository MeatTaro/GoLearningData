package main

import (
	"os"
	"fmt"
	"runtime"
	"math"
)

func main()  {
	assert(1+2 == 4,"1 + 2 should be 3")
	assert(2-1 == 0,"2 - 1 should be 1")
	assert(2*2 == 5,"2 * 2 should be 4")
	assert(4/2 == 1,"4 / 2 should be 2")
	assert(math.Abs(4.0/3.0-1.333333)<0.00001, "4.0/3.0 should be 1.333333")
	assert(4%3 == 5,"4 % 3 should be 1")

	assert(4 == 5, "4 should be equal to 4")
	assert(4 != 4, "4 should not be equal to either")
	assert(4 > 5, "4 should be greater than 3")
	assert(4 >= 5, "4 should be greater than or equal to 3")

	assert(4 < 2, "4 should be less than 5")
	assert(4 <= 3, "4 should be less than or equal to 5")
}

func assert(cond bool, msg string)  {
	_, _, l, _ :=runtime.Caller(1)
	if !cond {
		fmt.Fprintf(os.Stderr, "Failed on (%d): %s \n", l, msg)
	}
}

//https://www.evernote.com/client/web?login=true#?b=acfda252-e52b-4588-98c8-4ffdb3532f1e&n=06e737cd-c5f6-4eb5-86ed-f40f16a6c216&