package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"os"
)

func readFile(filename string) (answer [][]byte, lists [][]byte, wrapflag bool) {

	filecontents, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	rowstring := ""
	place := 0
	//read until we encounter " "
	for filecontents[place] != byte(32) {
		rowstring += string(filecontents[place])
		place++
	}

	place += 1 //skip space

	colstring := ""
	//read until we find newline character
	for filecontents[place] != byte(10) {
		colstring += string(filecontents[place])
		place++
	}

	place += 1

	numColumns, err := strconv.Atoi(colstring)
	if err != nil {
		panic(err)
	}
	numRows, _ := strconv.Atoi(rowstring)
	if err != nil {
		panic(err)
	}

	//initialize a 2-d slice 
	answer = make([][]byte, numRows)

	//read values in from file.
	for row := 0; row < numRows; row++ {
		answer[row] = make([]byte, numColumns)

		for col := 0; col < numColumns; col++ {
			answer[row][col] = filecontents[place+(1+numColumns)*row+col]
		}
	}

	//Print grid for debugging purposes
	/*
		for _, row := range answer {
			for _, char := range row {
				fmt.Print(string(char))
			}
			fmt.Println("")

		}
	*/

	place += (numColumns + 1) * numRows //skip reading past grid
	tempstring := ""
	//read until we find newline character
	for i := place; filecontents[i] != byte(10); i++ {
		tempstring += string(filecontents[i])
		place = i
	}

	if tempstring == "WRAP" {
		wrapflag = true
	} else if tempstring == "NO_WRAP" {
		wrapflag = false
	} else {
		panic("bad flag" + tempstring)
	}

	place += 2
	tempstring = ""
	//read until we find newline character
	for i := place; filecontents[i] != byte(10); i++ {
		tempstring += string(filecontents[i])
		place = i
	}

	numberoflists, err := strconv.Atoi(tempstring)
	if err != nil {
		panic(err)
	}

	lists = make([][]byte, numberoflists)

	for j := 0; j < numberoflists; j++ {

		place += 2
		tempstring = ""
		//read until we find newline character
		for i := place; i < len(filecontents) && filecontents[i] != byte(10); i++ {
			tempstring += string(filecontents[i])
			place = i
		}
		lists[j] = wordByte(tempstring)
	}

	return
}

func match(grid [][]byte, word []byte, xstart int, ystart int,
	xdir int, ydir int) bool {

	wordlength := len(word)

	xsize := len(grid)
	ysize := len(grid[0])

	//Check if we fall off grid
	if (wordlength-1)*xdir+xstart >= xsize || (wordlength-1)*ydir+ystart >= ysize ||
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

func dumbSearch(grid [][]byte, word []byte) (
	xstart int, ystart int, xend int, yend int) {
	numrows := len(grid)
	numcols := len(grid[0])

	xstart = -1

	for i := 0; i < numrows; i++ {
		for j := 0; j < numcols; j++ {

			for a := -1; a < 2; a++ {
				for b := -1; b < 2; b++ {
					if match(grid, word, i, j, a, b) {
						xstart = i
						ystart = j
						xend = i + a*(len(word)-1)
						yend = j + b*(len(word)-1)
						return
					}
				}
			}

		}
	}
	return
}

func main() {

//default input is test.txt
	inputname := "test.txt"

//unless it's specified
if len(os.Args) > 1{
	inputname = os.Args[1]
}


	grid, list, _ := readFile(inputname)

	for i := 0; i < len(list); i++ {

		x0, y0, x1, y1 := dumbSearch(grid, list[i])
		if x0 == -1 {
			fmt.Println("NOT FOUND")
		} else {
			fmt.Println("(", x0, ",", y0, ") (", x1, ",", y1, ")")
		}
	}

}
