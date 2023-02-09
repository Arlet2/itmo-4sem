package main

import (
	"fmt"
	. "lab1/internal"
)

func main() {
	fmt.Print("Выберите формат ввода 1) консоль, 2) файл: ")

	var mode int

	fmt.Scanln(&mode)

	var reader MatrixReader

	if mode == 1{
		reader = ConsoleReader{}
	} else if mode == 2{
		reader = FileReader{}
	} else {
		fmt.Println("Выбран некорректный режим. Завершение...")
		return
	}

	eps, matrix, err := reader.Read()

	if err != nil{
		fmt.Println("Ошибка: "+err.Error())
		return
	}

	fmt.Println("Eps:", eps)
	matrix.Print()

}