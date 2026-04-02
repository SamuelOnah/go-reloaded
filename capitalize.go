package main

import (
	"fmt"
	"strings"
)

func capitalizeFirstWord(samuelonah string) string {
	sam := strings.ToLower(samuelonah)
	return strings.Title(sam)
}

func main() {
	fmt.Println(capitalizeFirstWord("samuel onah"))
	fmt.Println(capitalizeFirstWord("oNAH SAMUEL"))
}

// this is a program that returns every first value as capital
// this you would notice in my project is that i first of all coverted all the strings to lower case which mean any input value will return as lowercase
// then later on in the code we told  it to return every first letter of lowercase into caplocks
