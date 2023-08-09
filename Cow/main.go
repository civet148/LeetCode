package main

import (
	"fmt"
)

const (
	COW_LIFE_YEARS = 6
)

var m_nCowCount int64

func main() {
	var years = 30
	fmt.Printf("after %v years alive cows [%v]\n", years, Cows(years))
}

// 一只牛寿命6岁（生下来0岁，6岁卒），每只牛3、5岁时可产小牛一只（寿命同样是6岁）
// 写一个方法，给出年限，得出活着的牛数量(初始牛数量为1)
func Cows(years int) (cows int) {
	left := years
	if left > COW_LIFE_YEARS {
		left = COW_LIFE_YEARS
	}
	cows++
	for i := 0; i <= left; i++ {
		if i == 3 || i == 5 {
			cows += Cows(years - i)
		}
		if i == COW_LIFE_YEARS {
			cows--
		}
	}
	return
}
