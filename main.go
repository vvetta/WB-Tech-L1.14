package main

import "fmt"

func main() {
	checkType("Привет")
	checkType(123)
	checkType(true)
	checkType(make(chan int))
	checkType(make(chan string))
}

func checkType(val interface{}) {
	switch val.(type) {
		case string:
			fmt.Println("Строка!")
		case int:
			fmt.Println("Число!")
		case chan int:
			fmt.Println("Канал int!")
		case chan string:
			fmt.Println("Канал string!")
		case bool:
			fmt.Println("Bool!")
		default:
			fmt.Println("Такого мы не проходили...")
	}
}
