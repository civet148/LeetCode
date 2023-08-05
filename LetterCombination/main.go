package main

import (
	"fmt"
	"strconv"
)

func main() {
	strDigits := "2"
	fmt.Printf("输入：%v 输出: %v \n", strDigits, LetterCombinations(strDigits))
}

var mapDigitLetters = map[int64]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

//按照手机英文九键输入法（2~9），输入数字组合，返回可能的组合
//例如"23" 对应组合["ad","ae","af","bd","be","bf","cd","ce","cf"]
func LetterCombinations(digits string) (results []string) {
	var lst []string
	for _, v := range digits {

		n, _ := strconv.ParseInt(string(v), 10, 64)
		lst = append(lst, mapDigitLetters[n])
	}

	fmt.Printf("lst = %v\n", lst)
	nTotal := len(lst)

	if nTotal == 1 {
		for _, v := range lst[0] {
			results = append(results, string(v))
		}

	} else {
		for i := 0; i < nTotal-1; i++ {
			for j := i + 1; j < nTotal; j++ {

				for _, v := range lst[i] {

					for _, vv := range lst[j] {

						results = append(results, string(v)+string(vv))
					}
				}
			}
		}
	}

	return
}
