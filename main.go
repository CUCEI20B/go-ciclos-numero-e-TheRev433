package main

import "fmt"

func main() {
	var lim int
	var e float64 = 0
	fmt.Scanln(&lim)

	for i := 0; i <= lim; i++ {
		var fact float64 = 1
		for j := 1; j <= i; j++ {
			fact *= float64(j)
		}
		e += float64((float64(1) / fact))
	}
	fmt.Println(e)
}
