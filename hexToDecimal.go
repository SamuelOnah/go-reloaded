package main

import (
	"fmt"
	"strconv"
)

func hexToDecimal(hexStr string) (int64, error) {
	return strconv.ParseInt(hexStr, 16, 64)
}

func main() {
	fmt.Println(hexToDecimal("1E"))
	fmt.Println(hexToDecimal("FF"))

}

// This my very first public project on GitHub

// basically my program is about converting hexToDecimal, using string conversion to convert hexToDecimal
// if you run this program the expected out would be 30nil
