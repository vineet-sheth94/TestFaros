package main

import (
	solutions "./solutions"
	student "./student"
	"github.com/01-edu/z01"
)

func main() {
	args := []int{z01.RandIntBetween(2, 20)}

	for i := 0; i < 20; i++ {
		args = append(args, z01.RandIntBetween(2, 20))
	}

	for _, v := range args {
		z01.Challenge("ActiveBits", student.ActiveBits, solutions.ActiveBits, v)
	}
}
