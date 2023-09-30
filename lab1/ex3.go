package main

import "fmt"

func main() {
	//Инициализация переменных
	var userinit8 uint8 = 1
	var userinit16 uint16 = 2
	var userinit64 int64 = -3
	var userautoinit = -4 //Такой вариант инициализации также возможен

	fmt.Println("Values: ", userinit8, userinit16, userinit64, userautoinit, "\n")

	//Краткая запись объявления переменной
	//только для новых переменных
	intVar := 10

	fmt.Printf("Value = %d Type = %T\n", intVar, intVar)

	//Задание.
	//1. Вывести типы всех переменных
	//2. Присвоить переменной intVar переменные userinit16 и userautoinit. Результат вывести.
	fmt.Printf("userinit8 = %d Type = %T\n", userinit8, userinit8)
	fmt.Printf("userinit16 = %d Type = %T\n", userinit16, userinit16)
	fmt.Printf("userinit64 = %d Type = %T\n", userinit64, userinit64)
	fmt.Printf("userautoinit = %d Type = %T\n", userautoinit, userautoinit)
	intVar = int(userinit16)
	fmt.Printf("intVar = %d Type = %T\n", intVar, intVar)
	intVar = userautoinit
	fmt.Printf("intVar = %d Type = %T\n", intVar, intVar)
}
