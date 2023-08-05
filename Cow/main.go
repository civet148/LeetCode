package main

import (
	"fmt"
	"time"
)

const (
	COW_LIFE_YEARS = 6
)

var m_nCowCount int64

func main() {
	var nYears = 30
	var nBeginTime = time.Now().Unix()
	m_nCowCount = GetAliveCows(nYears)
	var nEndTime = time.Now().Unix()
	fmt.Printf("%v年后活着的牛数量：[%v]只  -  计算耗时:%v 秒\n", nYears, m_nCowCount, nEndTime-nBeginTime)
}

// 一只牛寿命6岁（生下来0岁，6岁卒），每只牛3、5岁时可产小牛一只（寿命同样是6岁）
// 写一个方法，给出年限，得出活着的牛数量(初始牛数量为1)
func GetAliveCows(nYears int) (nCount int64) {
	return Cow(nYears)
}

func Cow(nYears int) (nCount int64) {
	nLeftYears := nYears
	if nLeftYears > COW_LIFE_YEARS {
		nLeftYears = COW_LIFE_YEARS
	}
	nCount++
	for i := 0; i <= nLeftYears; i++ {
		if i == 3 || i == 5 {
			nCount += Cow(nYears - i)
		} else if i == COW_LIFE_YEARS {
			nCount--
		}
	}
	return
}
