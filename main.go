package main

import (
	"fmt"
	"math/rand"
	"time"
)

var alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func fulfillFreeSpaces(matriz *[12][12]string) {
	for i := 0; i < len(*matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			if matriz[i][j] == "[ ]" {
				rand.Seed(time.Now().UnixNano())
				randomIndex := rand.Intn(26)
				matriz[i][j] = "[" + string(alphabet[randomIndex]) + "]"
			}
		}
	}
}

//func verifyCommonLetter(letter string, word string) bool {}

func setLocationAxisX(matriz *[12][12]string, word string) {
	rand.Seed(time.Now().UnixNano())
	randomRow := rand.Intn(12)
	randomColumn := rand.Intn(13 - len(word)) // subtract the word size to avoid false printing
	position := 0                             // set up an iterator for the word parameter

	// testing spaces
	for i := 0; i < len(word); i++ {
		//commonLetter := verifyCommonLetter(matriz[randomRow][i+randomColumn], word)
		if matriz[randomRow][i+randomColumn] != "[ ]" {
			println("!!!!!!!!One of the spaces is already filled, trying again!!!!!!!!")
			setLocationAxisX(matriz, word)
			return
		}
	}

	// fulfillment spaces
	for i := randomColumn; i < (randomColumn + len(word)); i++ {
		matriz[randomRow][i] = "[" + string(word[position]) + "]"
		position++
	}
}

// Similar to the setLocationAxisX function, set rows instead of columns as the main parameter
func setLocationAxisY(matriz *[12][12]string, word string) {
	rand.Seed(time.Now().UnixNano())
	randomColumn := rand.Intn(12)
	randomRow := rand.Intn(13 - len(word))
	position := 0

	for i := 0; i < len(word); i++ {
		//commonLetter := verifyCommonLetter(matriz[randomRow][i+randomColumn], word)
		if matriz[i+randomRow][randomColumn] != "[ ]" {
			println("!!!!!!!!One of the spaces is already filled, trying again!!!!!!!!")
			setLocationAxisY(matriz, word)
			return
		}
	}

	for i := randomRow; i < (randomRow + len(word)); i++ {
		matriz[i][randomColumn] = "[" + string(word[position]) + "]"
		position++
	}
}

func main() {

	wordsX := []string{"BIRD", "FISH", "TREE", "FOUR", "LION"}
	wordsY := []string{"RAIN", "BOOK", "LOVE", "MOON", "SONG"}

	// make sure to adjust the matrix size in the function parameters above accordingly
	matriz := [12][12]string{}

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			matriz[i][j] = "[ ]"
		}
	}

	for _, word := range wordsY {
		setLocationAxisY(&matriz, word)
	}

	for _, word := range wordsX {
		setLocationAxisX(&matriz, word)
	}

	fulfillFreeSpaces(&matriz)

	println("----------- FINAL RESULT -----------")
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			fmt.Print(matriz[i][j])
		}
		fmt.Print("\n")
	}
}
