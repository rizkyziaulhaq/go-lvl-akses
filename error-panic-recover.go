package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func tryError() {
	var input string
	fmt.Print("Type some number:")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input) // mengembalikan 2 data yg ditampung number dan err

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}
}

// membuat custom error
func validasi(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("tidak boleh kosong") // errors.New() bisa diganti dengan -> fmt.Errorf()
	}
	return true, nil
}

// recover
func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}
