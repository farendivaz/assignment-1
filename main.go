package main

import (
	"001/classmates"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please Enter Valid Absen")
		return
	}
	absen, err := strconv.Atoi(os.Args[1])
	_ = err

	if absen > len(classmates.Mentee) || absen < 1 {
		fmt.Println("Invalid Absen")
		return
	}

	classmates.Mentee[absen-1].Result()
}