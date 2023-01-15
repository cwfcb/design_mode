package adapter

import "fmt"

type computer interface {
	insertInSquarePort()
}

type mac struct {
}

func (m *mac) insertInSquarePort() {
	fmt.Println("Insert square port into mac machine")
}

type windows struct{}

func (w *windows) insertInCirclePort() {
	fmt.Println("Insert circle port into windows machine")
}

type client struct {
}

func (c *client) insertSquareUsbInComputer(com computer) {
	com.insertInSquarePort()
}
