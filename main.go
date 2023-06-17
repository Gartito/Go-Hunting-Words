package main

import (
	"fmt"
	"math/rand"
	"time"
)

var alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func fulfillFreeSpaces(matriz *[7][7]string) {
	for i := 0; i < len(*matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			if matriz[i][j] == "[ ]" {
				rand.Seed(time.Now().UnixNano())
				randomIndex := rand.Intn(27)
				matriz[i][j] = "[" + string(alphabet[randomIndex]) + "]"
			}
		}
	}
}

func setLocationAxisX(matriz *[7][7]string, word string) {
	rand.Seed(time.Now().UnixNano())
	randomRowPosition := rand.Intn(7)
	randomColumnPosition := rand.Intn(8 - len(word)) // subtract the word size to avoid false printing
	position := 0                                    // set up an iterator for the word parameter

	// testing spaces
	for i := 0; i < len(word); i++ {
		if matriz[randomRowPosition][i+randomColumnPosition] != "[ ]" {
			if string(word[i]) == string(matriz[randomRowPosition][i+randomColumnPosition][1]) {
				continue
			}
			setLocationAxisX(matriz, word)
			return
		}
	}

	// fulfillment spaces
	for i := randomColumnPosition; i < (randomColumnPosition + len(word)); i++ {
		matriz[randomRowPosition][i] = "[" + string(word[position]) + "]"
		position++
	}
}

// Similar to the setLocationAxisX function, set rows instead of columns as the main parameter
func setLocationAxisY(matriz *[7][7]string, word string) {
	rand.Seed(time.Now().UnixNano())
	randomColumnPosition := rand.Intn(7)
	randomRowPosition := rand.Intn(8 - len(word))
	position := 0

	for i := 0; i < len(word); i++ {
		if matriz[i+randomRowPosition][randomColumnPosition] != "[ ]" {
			if string(word[i]) == string(matriz[i+randomRowPosition][randomColumnPosition][1]) {
				continue
			}
			setLocationAxisX(matriz, word)
			return
		}
	}

	for i := randomRowPosition; i < (randomRowPosition + len(word)); i++ {
		matriz[i][randomColumnPosition] = "[" + string(word[position]) + "]"
		position++
	}
}

func main() {

	//wordsX := []string{"BIRD", "FISH", "TREE", "FOUR", "LION"}
	//wordsY := []string{"RAIN", "BOOK", "LOVE", "MOON", "SONG"}
	wordsX := []string{"MACACO", "ARROZ", "BOLO", "OVO"}
	wordsY := []string{"JOGO", "FACA", "TATOO"}

	// make sure to adjust the matrix size in the function parameters above accordingly
	matriz := [7][7]string{}

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

	//fulfillFreeSpaces(&matriz)

	println("----------- FINAL RESULT -----------")
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			fmt.Print(matriz[i][j])
		}
		fmt.Print("\n")
	}
}
