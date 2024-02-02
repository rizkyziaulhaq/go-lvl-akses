package library

import (
	"fmt"
)

func SayHello(name string) {
	fmt.Println("hello")
	introduce(name)
}

func introduce(name string) {
	fmt.Println("nama saya", name)
}

type Maba struct {
	Name string
	Age  int
}

func Mabaru() {
	var mahasiswa = Maba{"Zia", 22}

	mhs := []Maba{
		{
			Name: "Rizky",
			Age:  21,
		},
		{
			Name: "Ziaul",
			Age:  20,
		},
		{
			Name: "Haq",
			Age:  23,
		},
	}

	fmt.Println("\nname :", mahasiswa.Name)
	fmt.Println("Age :", mahasiswa.Age)
	fmt.Printf("\n%v", mhs)
}

var Student = struct {
	Name  string
	Grade int
}{}

func Init() {
	Student.Name = "Gojo Satoru"
	Student.Grade = 3

	fmt.Println("\n--> library/library.go imported")
	fmt.Printf("Name : %s\n", Student.Name)
	fmt.Printf("Grade : %d\n", Student.Grade)
}

func Perulangan() {
	counter := 0
	for counter <= 10 {
		counter++
		if counter%2 == 0 {
			fmt.Println("nilai counter bil.genap: ", counter)
		}
	}
	fmt.Println("diluar perulangan")
}
