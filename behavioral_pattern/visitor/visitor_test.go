package visitor

import "testing"

func TestVisitor(t *testing.T) {
	square := &square{side: 2}
	circle := &circle{radius: 2}
	rectangle := &rectangle{
		length: 2,
		width:  3,
	}

	var visitor visitor

	visitor = &areaCalculator{}
	square.accept(visitor)
	circle.accept(visitor)
	rectangle.accept(visitor)

	visitor = &middleCoordinates{}
	square.accept(visitor)
	circle.accept(visitor)
	rectangle.accept(visitor)
}
