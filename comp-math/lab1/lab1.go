package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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
	n, _ := fmt.Scanf("%d", &size)

	if n != 1 {
		fmt.Println("n - целое число")
		return 0, matrix{}, true
	}

	var eps float64

	fmt.Print("Введите точность: ")
	n, _ = fmt.Scanf("%f", &eps)

	if n != 1 {
		fmt.Println("Точность - число")
		return 0, matrix{}, true
	}

	coeff := make([][]int, size)

	fmt.Println("Вводите коэффициенты, разделяя строки переносом строки")
	for i:=0;i<size;i++{
		coeff[i] = make([]int, size+1)
		for j:=0;j<size+1;j++{
			n, _ = fmt.Scanf("%d", &coeff[i][j])

			if n != 1 {
				fmt.Println("Коэффициенты - целые числа")
				return 0, matrix{}, true
			}
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

	line, _, err := reader.ReadLine()

	if err == io.EOF {
		fmt.Println("Файл пустой")
		return 0, matrix{}, true
	}

	if err != nil {
		fmt.Println(err)
		return 0, matrix{}, true
	}

	size, err := strconv.Atoi(strings.Split(string(line), " ")[0])

	if err != nil {
		fmt.Println("Размер матрицы - целое число")
		return 0, matrix{}, true
	}

	eps, err := strconv.ParseFloat(strings.Split(string(line), " ")[1], 64)

	if err != nil {
		fmt.Println("Точность - число")
		return 0, matrix{}, true
	}

	coeff := make([][]int, size)

	for i:=0;i<size;i++{
		coeff[i] = make([]int, size+1)
		line, _, err = reader.ReadLine()

		if err != nil{
			fmt.Println("Недостаточно строк")
			return 0, matrix{}, true
		}

		if len(strings.Split(string(line), " ")) < size {
			fmt.Println("Недостаточно коэффициентов в строке")
			return 0, matrix{}, true
		}
		
		for j:=0;j<size+1;j++{
			n, _ = fmt.Scanf("%d", &coeff[i][j])

			if n != 1 {
				fmt.Println("Коэффициенты - целые числа")
				return 0, matrix{}, true
			}
		}
	}

	return eps, matrix{size: size, coeff: coeff}, false

}