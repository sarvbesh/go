package main

import (
	"fmt"
	"maps"
)

type Vertex struct {
	Lat, Long string
}

type Vertex1 struct {
	Lat, Long float64
}

var m1 = map[string]Vertex1{
	"Bell Labs": Vertex1{
		40.68433, -74.39967,
	},
	"Google": Vertex1{
		37.42202, -122.08408,
	},
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Headline"] = Vertex{
		"Hello, world!" , "My name is Sarvesh",
	}
	fmt.Println(m["Headline"])

	fmt.Println(m1)

	// empty array (map[key-type]value-type)
    m := make(map[string]int)

	// name[key] = value
    m["k1"] = 7
    m["k2"] = 13

	// show all of its key/value pairs
    fmt.Println("map:", m)

	// get the value of key name[key]
    v1 := m["k1"]
    fmt.Println("v1:", v1)

	// if key doesn't exist, zero value is returned
    v3 := m["k3"]
    fmt.Println("v3:", v3)

	// returns number of key/value pairs
    fmt.Println("len:", len(m))

	// removes key/value pairs
    delete(m, "k2")
    fmt.Println("map:", m)

	// removes all key/value pairs
    clear(m)
    fmt.Println("map:", m)

	/*
	optional second return value when getting a value from a map indicates if the key was present in the map 
	this can be used to disambiguate between missing keys and keys with zero values like 0 or "" 
	here we didn’t need the value itself, so we ignored it with the blank identifier _
	
	*/
    _, prs := m["k2"]
    fmt.Println("prs:", prs)

	// declare and initialize in same line
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

    n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }
	
	// more examples ~ mutating maps

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
