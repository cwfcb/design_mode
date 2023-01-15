package interpreter

import (
	"errors"
	"strings"
)

/*
	运算符只包含加、减、乘、除，并且没有优先级的概念；
	表达式（也就是前面提到的“句子”）中，先书写数字，后书写运算符，空格隔开；
	按照先后顺序，取出两个数字和一个运算符计算结果，结果重新放入数字的最头部位置，循环上述过程，直到只剩下一个数字，这个数字就是表达式最终的计算结果。
*/

type ExpressionInterpreter struct {
}

func (e *ExpressionInterpreter) interpreter(expression string) (int, error) {
	elements := strings.Split(expression, " ")
	var numExpression []Expression

	for i := 0; i < (len(elements)+1)>>1; i++ {
		numExpression = append(numExpression, NewNumberExpression(elements[i]))
	}

	var exp Expression
	for i := (len(elements) + 1) >> 1; i < len(elements); i++ {
		exp1, exp2 := numExpression[len(numExpression)-1], numExpression[len(numExpression)-2]
		numExpression = numExpression[:len(numExpression)-2]
		switch elements[i] {
		case "+":
			exp = NewAdditionExpression(exp1, exp2)
		case "-":
			exp = NewSubtractionExpression(exp1, exp2)
		case "*":
			exp = NewMultiplicationExpression(exp1, exp2)
		case "/":
			exp = NewDivisionExpression(exp1, exp2)
		default:
			return 0, errors.New("invalid operator")
		}
		numExpression = append(numExpression, exp)
	}
	return numExpression[0].interpreter(), nil
}
