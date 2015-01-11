package main

import "fmt"

type Display interface {
	open()
	print()
	close()
}

type AbstractDisplay struct {
	display Display
}
func newAbstractDisplay(d Display) AbstractDisplay {
	return AbstractDisplay{display: d}
}

func (a AbstractDisplay) Display() {
	a.display.open()
	for i := 0; i < 5; i++ {
		a.display.print()
	}
	a.display.close()
}

type CharDisplay struct {

}

func NewCharDisplay() AbstractDisplay {
	return newAbstractDisplay(CharDisplay{})
}

func (d CharDisplay) open() {
	fmt.Println("---------------")
}
func (d CharDisplay) print() {
	fmt.Print("d")
}
func (d CharDisplay) close() {
	fmt.Println("\n---------------")
}

var _ Display = (*CharDisplay)(nil)

func main() {
	d := NewCharDisplay()
	d.Display()
}

