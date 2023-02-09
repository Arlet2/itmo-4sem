package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type matrix struct {
	size int
	coeff [][]int
}

func (m matrix) print() {
	fmt.Printf("Matrix %dx%d:\n", m.size, m.size)

	for i:=0;i<m.size;i++{
		for j:=0;j<m.size+1;j++{
			fmt.Print(m.coeff[i][j], " ")
		}
		fmt.Println()
	}
}

type MatrixReader interface{
	read() (float64, matrix, bool)
}

type ConsoleReader struct {}
type FileReader struct {}

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

	eps, matrix, isErr := reader.read()

	if isErr{
		return
	}

	fmt.Println("Eps:", eps)
	matrix.print()

}

func (_ ConsoleReader) read() (float64, matrix, bool) {

	var size int

	fmt.Print("Введите размер матрицы: ")
	fmt.Scanln(&size) // проверка на что-то плохое

	var eps float64

	fmt.Print("Введите точность: ")
	fmt.Scanln(&eps) // проверка на что-то плохое

	coeff := make([][]int, size)

	fmt.Println("Вводите коэффициенты, разделяя строки переносом строки")
	for i:=0;i<size;i++{
		coeff[i] = make([]int, size+1)
		for j:=0;j<size+1;j++{
			fmt.Scan(&coeff[i][j]) // проверка на что-то плохое
		}
	}

	return eps, matrix{size: size, coeff: coeff}, false
}

func (_ FileReader) read() (float64, matrix, bool) {
	var path string
	fmt.Print("Введите название файла: ")

	fmt.Scanln(&path)

	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeIrregular)

	if err != nil{
		fmt.Println(err)
		return 0,matrix{}, true
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var size int

	line, _, err := reader.ReadLine()

	if err == io.EOF {
		fmt.Println("Файл пустой")
		return 0, matrix{}, true
	}

	if err != nil {
		fmt.Println(err)
		return 0, matrix{}, true
	}

	size = strings.to

	return 0, matrix{}, true

}