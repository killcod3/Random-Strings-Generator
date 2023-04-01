package main

import (
	"fmt"
	"os"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
)

var totalCombinations int
var combinationsGenerated int

func generateStrings(pattern string, index int, currentString string, file *os.File) {
	if index == len(pattern) {
		file.WriteString(currentString + "\n")
		combinationsGenerated++
		if combinationsGenerated%1000 == 0 {
			fmt.Printf("\r%.2f%%", float64(combinationsGenerated)/float64(totalCombinations)*100)
		}
		return
	}

	switch pattern[index] {
	case 'L':
		for _, char := range lowercase {
			generateStrings(pattern, index+1, currentString+string(char), file)
		}
	case 'U':
		for _, char := range uppercase {
			generateStrings(pattern, index+1, currentString+string(char), file)
		}
	case 'D':
		for _, char := range digits {
			generateStrings(pattern, index+1, currentString+string(char), file)
		}
	case 'R':
		allChars := lowercase + uppercase + digits
		for _, char := range allChars {
			generateStrings(pattern, index+1, currentString+string(char), file)
		}
	default:
		generateStrings(pattern, index+1, currentString+string(pattern[index]), file)
	}
}

func calculateTotalCombinations(pattern string) int {
	totalCombinations := 1
	for _, char := range pattern {
		switch char {
		case 'L':
			totalCombinations *= len(lowercase)
		case 'U':
			totalCombinations *= len(uppercase)
		case 'D':
			totalCombinations *= len(digits)
		case 'R':
			totalCombinations *= len(lowercase) + len(uppercase) + len(digits)
		default:
			totalCombinations *= 1
		}
	}
	return totalCombinations
}

func main() {
	pattern := "L-L" // Change this to the desired pattern

	file, err := os.Create("random_strings.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	totalCombinations = calculateTotalCombinations(pattern)

	generateStrings(pattern, 0, "", file)

	fmt.Println("\nDone!")
}