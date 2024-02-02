package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

// pengaksesan informasi property var objek
func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	reflectType := reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama		:", reflectType.Field(i).Name)
		fmt.Println("tipe data	:", reflectType.Field(i).Type)
		fmt.Println("nilai		:", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}
