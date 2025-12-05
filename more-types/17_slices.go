package main

import (
	"fmt"
	"slices"
)

/* 	gives more powerful interface to sequences than arrays

slices are typed only by the elements they contain (not the number of elements)

dynamically-sized, flexible view into the elements of an array

*/

func main() {

	// uninitialized slice equals to nil and has length 0

    var s []string
    fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// to create a slice with non-zero length, use builtin `make`

    s = make([]string, 3) // `string`s of length `3` (initially zero-valued)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) // by default a new slice's capacity is equal to its length

	// can set and get like arrays

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

	// returns length of slice

    fmt.Println("len:", len(s))

	// returns a slice containing one or more values

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

	// copy slice to an empty slice

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

	// slice operator ~ slice[low:high]

    l := s[2:5] // gets s[2], s[3], s[4]
    fmt.Println("sl1:", l)

	// slices up to (excluding) s[5]

    l = s[:5]
    fmt.Println("sl2:", l)

	// slices from (including) s[2]

    l = s[2:]
    fmt.Println("sl3:", l)

	// declare and initialize in a single line

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var xy []int = primes[1:4]
	fmt.Println(xy)

	// compare two slices

    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }

	// multidimensional slices

    twoD := make([][]int, 3)
    for i := range 3 {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := range innerLen {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}