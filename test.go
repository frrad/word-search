package main

import (
	"fmt"
	"math/rand"
	"time"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func ifneg(n int) int {
	if n < 0 {
		return -n
	}
	return 0
}

func main() {

	rand.Seed(time.Now().UnixNano())

	numberwords := 5

	rows := 15
	columns := 20
	//maxdimension := 0

	fmt.Println(rows, columns)

	/*	if rows > columns {
			maxdimension = rows
		} else {
			maxdimension = columns
		}*/

	table := make([][]string, rows)

	for j := 0; j < rows; j++ {
		table[j] = make([]string, columns)
		for i := 0; i < columns; i++ {
			table[j][i] = string(65 + rand.Int()%26)
		}

	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Print(table[i][j])
		}
		fmt.Println("")
	}

	fmt.Println("NO_WRAP")

	for i := 0; i < numberwords; i++ {

		const longestword = 4
		const shortestword = 2

		wordlength := rand.Int()%(longestword-shortestword+1) + shortestword

		a := rand.Int()%3 - 1
		b := rand.Int()%3 - 1
		if a == 0 && b == 0 {
			a = 1
		}

		x := rand.Int()%(rows-2*wordlength-2) + wordlength
		y := rand.Int()%(columns-2*wordlength-2) + wordlength

		word := ""
		for j := 0; j < wordlength; j++ {
			word += table[j*a+x][j*b+y]
		}
		fmt.Println(word)

	}

}
