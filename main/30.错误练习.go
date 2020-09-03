package main

import (
	"encoding/json"
	"fmt"
)
type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string{

	return 	fmt.Sprintf("cannot Sqrt negative number: %d",float64(e))

}

func MySqrt(x float64) (float64, error) {
	if x<0 {
		return 0,json.Unmarshal([]byte(""), 1)
	}
	z:=1.0
	for i:=0;i<20;i++ {
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}
	return z, nil
}

func main() {

	var z = ErrNegativeSqrt(2)

	fmt.Println()
	if s ,err := MySqrt(-2); err != nil {
		fmt.Println(s)
		fmt.Println(err)
	}
}
