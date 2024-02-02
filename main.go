package main

import (
	"fmt"
	"golang-level-akses/library"
	"reflect"
	"runtime"
)

// goroutine
// func print(till int, message string) {
// 	for i := 0; i < till; i++ {
// 		fmt.Println((i + 1), message)
// 	}
// }

func main() {

	// folder library
	library.SayHello("Zia")
	library.Mabaru()
	library.Init()
	library.Perulangan()

	// bangunDatar()     // interface

	// casting var interface kosong ke objek pointer
	type person struct {
		Name string
		Age  int
	}

	var secret interface{} = &person{Name: "john", Age: 30}
	var name = secret.(*person).Name
	// var age = secret.(*person).Age
	fmt.Println(name)

	//reflect packgae
	// var number = 21
	// var reflectValue := reflect.ValueOf(number)

	// fmt.Println("tipe variabel: ", reflectValue.Type())

	// if reflectValue.Kind() == reflect.Int {
	// 	fmt.Println("nilai variabel: ", reflectValue.Int()) // sesuai tipe data yg dicari example: Int, Float64, String
	// }

	// pengaksesan informasi property var objek
	var siswa = &student{Name: "Rachi", Grade: 3}
	siswa.getPropertyInfo()

	// pengaksesan informasi method var objek
	var s1 = &student{Name: "john wick", Grade: 2}
	fmt.Println("nama :", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("nama :", s1.Name)

	// goroutine
	// runtime.GOMAXPROCS(2)

	// go print(5, "halo")
	// print(5, "rachi")

	// var input string
	// fmt.Scanln(&input)

	// channel
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)

	// channel - select
	runtime.GOMAXPROCS(2)

	var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("\nnumbers: ", numbers)

	ch1 := make(chan float64)
	go getAverage(numbers, ch1)

	ch2 := make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Average \t %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t %d \n", max)
		}
	}

	// for - range - close pada channel
	runtime.GOMAXPROCS(2)

	var pesan = make(chan string)
	go sendMessage(pesan)
	printMessage(pesan)

	// channel timeout
	// runtime.GOMAXPROCS(2)

	// var attention = make(chan int)
	// go sendData(attention)
	// retreiveData(attention)

	// defer
	// defer fmt.Println("halo")
	// fmt.Println("welcome")

	orderSomeFood("Mie gacoan lvl 8")
	orderSomeFood("Piscok")
	nomerAntrian()

	// exit
	// keluar()

	// error
	tryError()

	// custom error
	defer catch() // penggunaan recover

	var nama string
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)

	if valid, err := validasi(nama); valid {
		fmt.Println("Hallo", nama)
	} else {
		// fmt.Println(err.Error())

		// penggunaan panic
		panic(err.Error())
	}

	// time
	waktu()
	parString()
	handlError()
	timeSleep()
	newTimer()
	afterFungsi()
	scheduler()
	tidur()
}
