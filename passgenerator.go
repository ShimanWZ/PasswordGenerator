package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	upperCase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerCase    = "abcdefghijklmnopqrstuvwxyz"
	numbers      = "0123456789"
	characters   = "@#!$%^&*_+=:><"
	charString   = upperCase + lowerCase + numbers + characters
	letterString = upperCase + lowerCase
)

func main() {
	var letterCount int
	var numberCount int
	var charCount int

	fmt.Print("Enter the number of letters you want in your password: ")
	fmt.Scanf("%d\n", &letterCount)
	fmt.Print("Enter the number of numbers you want in your password: ")
	fmt.Scanf("%d\n", &numberCount)
	fmt.Print("Enter the number of special characters in your password: ")
	fmt.Scanf("%d\n", &charCount)

	fmt.Printf("Your generated password is: %v", passGenerator(letterCount, numberCount, charCount))
}

func passGenerator(letterCount, numberCount, charCount int) string {
	var passString string = ""
	rand.Seed(time.Now().UnixNano())

	passString = passString + chooseSubStrings(letterCount, letterString)
	passString = passString + chooseSubStrings(charCount, characters)
	passString = passString + chooseSubStrings(numberCount, numbers)

	rand.Seed(time.Now().UnixNano())

	runeArr := []rune(passString)
	rand.Shuffle(len(runeArr), func(i, j int) { runeArr[i], runeArr[j] = runeArr[j], runeArr[i] })

	return string(runeArr)
}

func chooseSubStrings(count int, array string) string {
	var out string = ""
	var randomInt int

	for i := count; i > 0; i-- {
		randomInt = rand.Intn(len(array))
		out = out + string(array[randomInt])
	}
	return out
}
