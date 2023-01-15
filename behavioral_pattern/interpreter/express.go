package interpreter

import "strconv"

type Expression interface {
	interpreter() int
}

type NumberExpression struct {
	number int
}

func NewNumberExpression(num string) *NumberExpression {
	number, _ := strconv.Atoi(num)
	return &NumberExpression{number}
}

func (n *NumberExpression) interpreter() int {
	return n.number
}

type AdditionExpression struct {
	exp1, exp2 Expression
}

func NewAdditionExpression(exp1, exp2 Expression) *AdditionExpression {
	return &AdditionExpression{exp1, exp2}
}

func (a *AdditionExpression) interpreter() int {
	return a.exp1.interpreter() + a.exp2.interpreter()
}

type SubtractionExpression struct {
	exp1, exp2 Expression
}

func NewSubtractionExpression(exp1, exp2 Expression) *SubtractionExpression {
	return &SubtractionExpression{exp1, exp2}
}

func (s *SubtractionExpression) interpreter() int {
	return s.exp1.interpreter() + s.exp2.interpreter()
}

type MultiplicationExpression struct {
	exp1, exp2 Expression
}

func NewMultiplicationExpression(exp1, exp2 Expression) *MultiplicationExpression {
	return &MultiplicationExpression{exp1, exp2}
}

func (m *MultiplicationExpression) interpreter() int {
	return m.exp1.interpreter() + m.exp2.interpreter()
}

type DivisionExpression struct {
	exp1, exp2 Expression
}

func NewDivisionExpression(exp1, exp2 Expression) *DivisionExpression {
	return &DivisionExpression{exp1, exp2}
}

func (d *DivisionExpression) interpreter() int {
	return d.exp1.interpreter() + d.exp2.interpreter()
}
