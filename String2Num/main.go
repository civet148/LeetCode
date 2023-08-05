package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("结果：%v\n", String2Number("-123456789"))
}

func String2Number(strNumbers string) (number int) {

	var bMinus bool

	strNumbers = strings.TrimSpace(strNumbers)
	if strNumbers == "" {
		return
	}

	strInput := strNumbers
	if string(strNumbers[0]) == "-" {
		strInput = strNumbers[1:]
		bMinus = true
	}

	nCount := len(strInput)

	for i := 0; i < nCount; i++ {

		v, _ := strconv.ParseInt(string(strInput[i]), 10, 32)
		number += int(v) * int(math.Pow10(nCount-i-1))
	}

	if bMinus {
		return -number
	}
	return
}
