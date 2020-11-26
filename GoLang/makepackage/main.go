package main

import (
	"fmt"
	"makepackage/process"
)

func main()  {
	state := process.NewStation("粗铣面型")
	cnc := process.CNC{state}
	fmt.Println(cnc.GetName(), cnc.StateNum(), cnc.Machine())
}