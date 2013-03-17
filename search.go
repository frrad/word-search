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

	fmt.Println(filecontents)

	answer = make([][]byte, numRows)
	for row := 0; row < numRows; row++ {
		answer[row] = make([]byte, numColumns)

		for col := 0; col < numColumns; col++ {
			answer[row][col] = filecontents[(1+numColumns)*row+col]
		}
	}

	fmt.Println(answer)

	return
}

func main() {
	fmt.Println("hi! world")

	readFile("test.txt", 3, 7)
}
