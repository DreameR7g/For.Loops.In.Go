package main

import (
	"fmt"
)

func main() {
	number := 10
	for i := 0; i < 100; i++ {
		number += i
	}
	fmt.Println(number)
	while()
}

// Output results in 4960 from main, 1024 from while

func while() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	return
}
