package bridge

import (
	"testing"
)

func TestComputer(t *testing.T) {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}
	macComputer := &mac{}
	macComputer.setPrinter(hpPrinter)
	macComputer.print()
	t.Log()
	macComputer.setPrinter(epsonPrinter)
	macComputer.print()
	t.Log()
	winComputer := &windows{}
	winComputer.setPrinter(hpPrinter)
	winComputer.print()
	t.Log()
	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
	t.Log()
}
