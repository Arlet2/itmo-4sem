package internal
import (
	"strconv"
	"strings"
	"bufio"
	"fmt"
	"io"
	"os"
)

type MatrixReader interface {
	Read() (float64, matrix, bool)
}

type ConsoleReader struct {}
type FileReader struct {}

func (ConsoleReader) Read() (float64, matrix, bool) {
	var input string

	fmt.Print("Введите размер матрицы: ")
	_, err := fmt.Scan(&input)
	fmt.Scanf(" ") // сброс пробела....

	if err != nil {
		panic("Нужно что-то ввести")
	}

	inputSize, err := strconv.ParseInt(input, 10, 8)

	if err != nil {
		panic("Размер матрицы - целое число")
	}

	if inputSize <= 0 {
		panic("Размер матрицы - положительное число")
	}

	var size int = int(inputSize)


	fmt.Print("Введите точность: ")
	_, err = fmt.Scan(&input)
	fmt.Scanf(" ") // сброс пробела....

	if err != nil {
		panic("Нужно что-то ввести")
	}

	eps, err := strconv.ParseFloat(input, 64)

	if err != nil {
		panic("Точность должна быть числом")
	}

	if eps <= 0 {
		panic("Точность должна быть положительным числом")
	}

	coeff := make([][]float64, size)

	fmt.Println("Вводите коэффициенты, разделяя строки переносом строки")
	for i := 0;i<size;i++{
		coeff[i] = make([]float64, size+1)
		for j:=0;j<size+1;j++{
			_, err = fmt.Scan(&input)

			if err != nil {
				panic("Недостаточно коэффициентов")
			}

			coeff[i][j], err = strconv.ParseFloat(input, 64)

			if err != nil {
				panic("Коэффициенты должны быть числами")
			}
		}
	}

	return eps, matrix{size: size, coeff: coeff}, false
}

func (FileReader) Read() (float64, matrix, bool) {
	var path string
	fmt.Print("Введите название файла: ")

	fmt.Scanln(&path)

	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeIrregular)

	if err != nil{
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	line, _, err := reader.ReadLine()

	if err == io.EOF {
		panic("Файл пустой")
	}

	if err != nil {
		panic(err)
	}

	size, err := strconv.Atoi(strings.Split(string(line), " ")[0])

	if err != nil {
		panic("Размер матрицы - целое число")
	}

	if size <= 0 {
		panic("Размер матрицы - положительное число")
	}

	eps, err := strconv.ParseFloat(strings.Split(string(line), " ")[1], 64)

	if err != nil {
		panic("Точность - число")
	}

	if eps <= 0 {
		panic("Точность - положительное число")
	}

	coeff := make([][]float64, size)

	for i:=0;i<size;i++{
		coeff[i] = make([]float64, size+1)
		line, _, err = reader.ReadLine()

		lineCoeff := strings.Split(string(line), " ")

		if err != nil{
			panic("Недостаточно строк")
		}

		if len(lineCoeff) < size {
			panic("Недостаточно коэффициентов в строке")
		}
		
		for j:=0;j<size+1;j++{

			value, err := strconv.ParseFloat(lineCoeff[j], 64)

			if err != nil {
				panic("Коэффициенты должны быть числами")
			}

			coeff[i][j] = value
		}
	}

	return eps, matrix{size: size, coeff: coeff}, false

}