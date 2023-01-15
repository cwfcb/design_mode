package visitor

import (
	"fmt"
	"math"
)

type visitor interface {
	// 定义多个抽象方法，是因为 golang 不支持方法重载，简单点说，就是不支持方法同名
	visitForSquare(*square)
	visitForCircle(*circle)
	visitFoRectangle(*rectangle)
}

type areaCalculator struct {
}

func (a *areaCalculator) visitForSquare(s *square) {
	//Calculate area for square. After calculating the area assign in to the area instance variable
	area := s.side * s.side
	fmt.Printf("Calculating area for square, side: %d, area: %d\n", s.side, area)
}

func (a *areaCalculator) visitForCircle(s *circle) {
	//Calculate are for circle. After calculating the area assign in to the area instance variable
	radius := float64(s.radius)
	area := math.Pi * radius * radius
	fmt.Printf("Calculating area for circle radisu: %.3f, area: %.3f\n", radius, area)
}

func (a *areaCalculator) visitFoRectangle(s *rectangle) {
	//Calculate are for rectangle. After calculating the area assign in to the area instance variable
	area := s.length * s.width
	fmt.Printf("Calculating area for rectangle, length: %d, width: %d, area: %d\n", s.length, s.width, area)
}

type middleCoordinates struct {
}

func (a *middleCoordinates) visitForSquare(s *square) {
	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *middleCoordinates) visitForCircle(c *circle) {
	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for circle")
}

func (a *middleCoordinates) visitFoRectangle(t *rectangle) {
	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for rectangle")
}
