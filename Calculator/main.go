package main

import (
	"fmt"
	"strings"
)

type Operator string

const (
	OperatorNull Operator = ""
	OperatorAdd  Operator = "+"
	OperatorSub  Operator = "-"
	OperatorMul  Operator = "*"
	OperatorDiv  Operator = "/"
)
const (
	LeftBracket  = "("
	RightBracket = ")"
)

type NodeExpr struct {
	Expr      string
	Operator  Operator
	LeftExpr  *NodeExpr
	RightExpr *NodeExpr
}

func main() {
	var strExpress = "2+(3*5/2+3-1)*2/2+8/4*(6-10)*2+100" //result=95.5
	node := NewNodeExpr(strExpress)
	fmt.Printf("计算结果：%v = %v \n", strExpress, node.Result())
}

func NewNodeExpr(strExpress string) *NodeExpr {
	strExpress = strings.TrimSpace(strExpress)
	return &NodeExpr{
		Expr:      strExpress,
		Operator:  OperatorNull,
		LeftExpr:  nil,
		RightExpr: nil,
	}
}

func (n *NodeExpr) Result() (result float64) {
	return
}

func findOperator(expr string) (idx int) {
	if expr == "" {
		return -1
	}
	idx = strings.Index(expr, string(OperatorAdd))
	if idx >= 0 {
		return idx
	}
	idx = strings.Index(expr, string(OperatorSub))
	if idx >= 0 {
		return idx
	}
	idx = strings.Index(expr, string(OperatorDiv))
	if idx >= 0 {
		return idx
	}
	idx = strings.Index(expr, string(OperatorMul))
	if idx >= 0 {
		return idx
	}
	return -1
}

func subExprs(expr string) (leftExpr, rightExpr string, operator Operator) {
	var lidx, ridx int
	if lidx = strings.Index(expr, LeftBracket); lidx != -1 { //左圆括号存在
		ridx = strings.Index(expr, RightBracket)
		if ridx == -1 {
			panic("right bracket not found") //语法错误: 左右圆括号不匹配
		}
	}
	return
}

func calculate(expr *NodeExpr) (r float64) {

	return
}

//
//
//func getSubExpress(strExpress string) (exp1, exp2, oper string, ok bool) {
//	var idx int
//	if idx = strings.Index(strExpress, "+"); idx != -1 {
//		ok = true
//		oper = "+"
//	} else if idx = strings.Index(strExpress, "-"); idx != -1 {
//		ok = true
//		oper = "-"
//	} else if idx = strings.Index(strExpress, "*"); idx != -1 {
//		ok = true
//		oper = "*"
//	} else if idx = strings.Index(strExpress, "/"); idx != -1 {
//		ok = true
//		oper = "/"
//	}
//
//	if ok {
//		exp1 = strExpress[:idx]
//		exp2 = strExpress[idx+1:]
//	}
//	//fmt.Printf("ok [%v] exp1 [%v] exp2 [%v] oper [%v] idx[%v] \n", ok, exp1, exp2, oper, idx)
//	return
//}
//
////输入计算公式字符串，得出计算结果
////符号: + - * /  字符：非负整数
//func Calculate(strExpress string) (nResult int64) {
//
//	if exp1, exp2, oper, ok := getSubExpress(strExpress); ok {
//		nResult = CalcTwoExpress(exp1, exp2, oper)
//	}
//	return
//}
//
//func CalcTwoExpress(exp1, exp2 string, oper string) (nResult int64) {
//
//	var r1, r2 int64
//	if e11, e12, oper1, ok := getSubExpress(exp1); ok {
//
//		r1 += CalcTwoExpress(e11, e12, oper1)
//	} else {
//		r1, _ = strconv.ParseInt(exp1, 10, 32)
//	}
//
//	if e21, e22, oper2, ok := getSubExpress(exp2); ok {
//
//		r2 += CalcTwoExpress(e21, e22, oper2)
//	} else {
//		r2, _ = strconv.ParseInt(exp2, 10, 32)
//	}
//
//	switch oper {
//	case "*":
//		nResult = r1 * r2
//	case "/":
//		nResult = r1 / r2
//	case "+":
//		nResult = r1 + r2
//	case "-":
//		nResult = r1 - r2
//	}
//	return
//}
