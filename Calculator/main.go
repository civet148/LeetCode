package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var strExpress = "2+3*2+8/4*6-10*2+100" //10

	fmt.Printf("计算结果：%v = %v \n", strExpress, Calculate(strExpress))

}

func getSubExpress(strExpress string) (exp1, exp2, oper string, ok bool) {
	var idx int
	if idx = strings.Index(strExpress, "+"); idx != -1 {
		ok = true
		oper = "+"
	} else if idx = strings.Index(strExpress, "-"); idx != -1 {
		ok = true
		oper = "-"
	} else if idx = strings.Index(strExpress, "*"); idx != -1 {
		ok = true
		oper = "*"
	} else if idx = strings.Index(strExpress, "/"); idx != -1 {
		ok = true
		oper = "/"
	}

	if ok {
		exp1 = strExpress[:idx]
		exp2 = strExpress[idx+1:]
	}
	//fmt.Printf("ok [%v] exp1 [%v] exp2 [%v] oper [%v] idx[%v] \n", ok, exp1, exp2, oper, idx)
	return
}

//输入计算公式字符串，得出计算结果
//符号: + - * /  字符：非负整数
func Calculate(strExpress string) (nResult int64) {

	if exp1, exp2, oper, ok := getSubExpress(strExpress); ok {
		nResult = CalcTwoExpress(exp1, exp2, oper)
	}
	return
}

func CalcTwoExpress(exp1, exp2 string, oper string) (nResult int64) {

	var r1, r2 int64
	if e11, e12, oper1, ok := getSubExpress(exp1); ok {

		r1 += CalcTwoExpress(e11, e12, oper1)
	} else {
		r1, _ = strconv.ParseInt(exp1, 10, 32)
	}

	if e21, e22, oper2, ok := getSubExpress(exp2); ok {

		r2 += CalcTwoExpress(e21, e22, oper2)
	} else {
		r2, _ = strconv.ParseInt(exp2, 10, 32)
	}

	switch oper {
	case "*":
		nResult = r1 * r2
	case "/":
		nResult = r1 / r2
	case "+":
		nResult = r1 + r2
	case "-":
		nResult = r1 - r2
	}
	return
}
