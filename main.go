package main

import (
	"math/rand/v2"
)

func main() {
	if rand.Float32() >= 0.5 {
		panic("panic!!!")
	} else {
		println("hellow world.")
	}
}
