package bridge

import "fmt"

type Computer interface {
	print()
	setPrinter(printer Printer)
}

type mac struct {
	printer Printer
}

func (m *mac) print() {
	fmt.Println("Print request for mac")
	m.printer.printFile()
}

func (m *mac) setPrinter(printer Printer) {
	m.printer = printer
}

type windows struct {
	printer Printer
}

func (w *windows) print() {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}

func (w *windows) setPrinter(printer Printer) {
	w.printer = printer
}
