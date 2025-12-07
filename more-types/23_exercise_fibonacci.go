package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int

func fibonacci() func() int {

	// store the previous two fibonacci numbers
    a, b := 0, 1
    
    // returned function is the closure
    return func() int {
        // store the current value of 'a' as the result to return.
        result := a
        
        // calculate the next pair:
        // new 'a' becomes the old 'b'
        // new 'b' becomes the sum (old 'a' + old 'b')
        a, b = b, a + b
        
        // return the stored result (the current Fibonacci number).
        return result
    }
}

func main() {
    
	// f is the closure, it has its own private 'a' and 'b' state
	f := fibonacci()
	
    fmt.Println("First 10 Fibonacci Numbers:")
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}