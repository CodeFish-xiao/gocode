package main

import (
	"fmt"
	"math"
)



func Sqrt(x float64) float64 {
	z:=1.0
	for i:=0;i<20;i++ {
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(3))

	fmt.Println(math.Sqrt(3))
}
