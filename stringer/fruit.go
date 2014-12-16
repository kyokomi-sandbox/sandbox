package stringer

import (
	"fmt"
)

// stringer -type Fruit fruit.go
type Fruit int

const (
	Apple Fruit = iota
	Orange
	Banana
)

func StringerExample() {
	var fruit Fruit = Apple
	fmt.Println(fruit)
}
