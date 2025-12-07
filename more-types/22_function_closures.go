package main

import "fmt"

/*
* it is a function that remembers and can access the variables from the scope where it was created,
* even after that original scope has finished executing
 */

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
