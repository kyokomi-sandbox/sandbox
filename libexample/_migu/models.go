package main

type User struct {
	ID int64 `migu:"pk"`
	Name string
	Age  uint
}
