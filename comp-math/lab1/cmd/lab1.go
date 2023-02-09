package main

import (
	"fmt"
	. "lab1/internal"
)

func main() {
	fmt.Print("Выберите формат ввода 1) консоль, 2) файл 3) lab1.txt: ")

	var mode int

	fmt.Scanln(&mode)

	var reader MatrixReader

	if mode == 1 {
		reader = ConsoleReader{}
	} else if mode == 2 {
		reader = FileReader{}
	} else if mode == 3 {
		reader = PreparedReader{Path: "lab1.txt"}
	} else {
		fmt.Println("Выбран некорректный режим. Завершение...")
		return
	}

	eps, matrix, err := reader.Read()

	if err != nil {
		fmt.Println("Ошибка: " + err.Error())
		return
	}

	fmt.Println("\nВведенные данные:")
	fmt.Println("Точность:", eps)
	matrix.PrintAugmented()

	if !matrix.TryToCreateDiagonalDominance() {
		matrix.PrintAugmented()
		fmt.Println("Эта матрица не обладает диагональным преобладанием. Применение метода невозможно")
		return
	}

	matrix.UseGaussZeidel(eps, false)

}
