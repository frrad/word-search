package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	shortestword    = 3
	minWordQuantity = 2
	maxWordQuantity = 10
	minxdim         = 10
	maxxdim         = 20
	minydim         = 10
	maxydim         = 20
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

func randInt(min int, max int) int {
	return rand.Int()%(max-min) + min
}
func main() {

	rand.Seed(time.Now().UnixNano())

	numberwords := randInt(maxWordQuantity, minWordQuantity)

	rows := rand.Int()%(maxxdim-minxdim) + minxdim
	columns := rand.Int()%(maxydim-minydim) + minydim

	fmt.Println(rows, columns)

	mindimension := 0
	if rows > columns {
		mindimension = columns
	} else {
		mindimension = rows
	}

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
	fmt.Println(numberwords)

	for i := 0; i < numberwords; i++ {

		longestword := mindimension

		wordlength := rand.Int()%(longestword-shortestword+1) + shortestword

		a := rand.Int()%3 - 1
		b := rand.Int()%3 - 1
		if a == 0 && b == 0 {
			a = 1
		}

		x := 0
		if a == 1 {
			x = rand.Int() % (rows - (wordlength - 1))
		} else if a == 0 {
			x = rand.Int() % (rows)
		} else {
			x = rand.Int()%(rows-(wordlength-1)) + (wordlength - 1)
		}

		y := 0
		if b == 1 {
			y = rand.Int() % (columns - (wordlength - 1))
		} else if b == 0 {
			y = rand.Int() % (columns)
		} else {
			y = rand.Int()%(columns-wordlength-1) + (wordlength - 1)
		}

		word := ""
		for j := 0; j < wordlength; j++ {
			word += table[j*a+x][j*b+y]
		}
		fmt.Println(word)

	}

}
