package main

import (
	"fmt"
	"strconv"
)

func binToDecimal(binStr string) (int64, error) {
	return strconv.ParseInt(binStr, 2, 64)
}

func main() {
	fmt.Println(binToDecimal("10"))
	fmt.Println(binToDecimal("1010"))
}

// this current program just has slight difference from that of hexToDecimal, this time around we are converting binary to decimal
// our expected output this time around is 1, and 10
