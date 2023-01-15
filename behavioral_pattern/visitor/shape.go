package visitor

type shape interface {
	getType() string
	accept(visitor)
}

type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

func (s *square) getType() string {
	return "Square"
}

type circle struct {
	radius int
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

func (c *circle) getType() string {
	return "Circle"
}

type rectangle struct {
	length int
	width  int
}

func (t *rectangle) accept(v visitor) {
	v.visitFoRectangle(t)
}

func (t *rectangle) getType() string {
	return "rectangle"
}
