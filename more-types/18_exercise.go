package main

import "golang.org/x/tour/pic"

// Pic should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers

func Pic(dx, dy int) [][]uint8 {
	
	// allocate the outer slice (rows) of length dy
	picture := make([][]uint8, dy) 

	// loop over the rows (y-coordinate)
	for y := 0; y < dy; y++ {
		// allocate the inner slice (columns) of length dx for current row
		picture[y] = make([]uint8, dx)

		// loop over the columns (x-coordinate)
		for x := 0; x < dx; x++ {
			// choose an interesting function for pixel value
			//	 - (x XOR y) * 8
			//   - (x + y) / 2
			//   - x * y
			//   - x ^ y
			
			// calculate the pixel value (an integer)
			intValue := (x ^ y) * 8
			
			// convert the integer value to uint8 and assign it to the pixel
			picture[y][x] = uint8(intValue)
		}
	}

	// return the complete picture data.
	return picture
}

func main() {
	pic.Show(Pic)
}