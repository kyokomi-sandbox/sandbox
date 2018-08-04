package main

import (
	"fmt"

	"math/rand"
	"time"

	"github.com/speps/go-hashids"
)

var ra = rand.New(rand.NewSource(time.Now().Unix()))

func randomID() int64 {
	return random(100000, 999999)
}

//var baseNumbers = []int64{901938, 661878, 655318}
var salt = "15e0b341-4e1d-4072-afb7-c1568af58262"

func random(min, max int64) int64 {
	return ra.Int63n(max-min) + min
}

func main() {
	hd := hashids.NewData()
	hd.Salt = salt
	h, _ := hashids.NewWithData(hd)
	id, _ := h.EncodeInt64([]int64{randomID(), randomID(), randomID()})
	numbers, _ := h.DecodeWithError(id)
	fmt.Println(id, numbers)
}
