package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "1234567890abcdefghijk_1234567890abcdefghijk_1234567890abcdefghijk"
	hashedPassword, hashError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashError != nil {
		panic(hashError.Error())
	}

	fmt.Println(string(hashedPassword), len(hashedPassword))

	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		fmt.Println("NG")
	} else {
		fmt.Println("OK")
	}
}
