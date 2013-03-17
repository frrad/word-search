package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(filename string, numRows int, numColumns int) (answer [][]byte) {

	filecontents, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	//initialize a 2-d slice 
	answer = make([][]byte, numRows)

	//read values in from file.
	for row := 0; row < numRows; row++ {
		answer[row] = make([]byte, numColumns)

		for col := 0; col < numColumns; col++ {
			answer[row][col] = filecontents[(1+numColumns)*row+col]
		}
	}

	//Print grid for debugging purposes
	/*
		for _, row := range answer {
			for _, char := range row {
				fmt.Print(char)
			}
			fmt.Println("")

		}
	*/

	return
}

func match(grid [][]byte, word []byte, xstart int, ystart int,
	xdir int, ydir int) bool {

	wordlength := len(word)

	xsize := len(grid)
	ysize := len(grid[0])

	//Check if we fall off grid
	if (wordlength-1)*xdir+xstart > xsize || (wordlength-1)*ydir+ystart > ysize ||
		(wordlength-1)*xdir+xstart < 0 || (wordlength-1)*ydir+ystart < 0 {
		return false
	}

	for i := 0; i < wordlength; i++ {
		if word[i] != grid[xstart+(i*xdir)][ystart+(i*ydir)] {
			return false
		}
	}
	return true

}

func wordByte(word string) (answer []byte) {

	answer = make([]byte, len(word))

	for i, letter := range word {
		answer[i] = byte(letter)
	}
	return
}

func main() {

	numrows := 3
	numcols := 7
	grid := readFile("test.txt", numrows, numcols)
	word := wordByte("TME")

	for i := 0; i < numrows; i++ {
		for j := 0; j < numcols; j++ {

			for a := -1; a < 2; a++ {
				for b := -1; b < 2; b++ {
					if match(grid, word, i, j, a, b) {
						fmt.Println("(", 1+i, ",", j+1, ")", a, b)
					}
				}
			}

		}
	}

}
