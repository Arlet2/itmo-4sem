package internal

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type MatrixReader interface {
	Read() (float64, matrix, error)
}

type ConsoleReader struct {}
type FileReader struct {}
type PreparedReader struct {
	FileReader
	Path string
}

func (ConsoleReader) Read() (float64, matrix, error) {
	var input string

	fmt.Print("Введите размер матрицы: ")
	_, err := fmt.Scan(&input)
	fmt.Scanf(" ") // сброс пробела....

	if err != nil {
		return 0, matrix{}, errors.New("нужно что-то ввести")
	}

	inputSize, err := strconv.ParseInt(input, 10, 8)

	if err != nil {
		return 0, matrix{}, errors.New("размер матрицы - целое число")
	}

	if inputSize <= 0 {
		return 0, matrix{}, errors.New("размер матрицы - положительное число")
	}

	var size int = int(inputSize)


	fmt.Print("Введите точность: ")
	_, err = fmt.Scan(&input)
	fmt.Scanf(" ") // сброс пробела....
	input = strings.ReplaceAll(input, ",", ".")

	if err != nil {
		return 0, matrix{}, errors.New("нужно что-то ввести")
	}

	eps, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, matrix{}, errors.New("точность должна быть числом")
	}

	if eps <= 0 {
		return 0, matrix{}, errors.New("точность должна быть положительным числом")
	}

	coeff := make([][]float64, size)

	fmt.Println("Вводите коэффициенты, разделяя строки переносом строки")
	for i := 0;i<size;i++{
		coeff[i] = make([]float64, size+1)
		for j:=0;j<size+1;j++{
			_, err = fmt.Scan(&input)

			if err != nil {
				return 0, matrix{}, errors.New("недостаточно коэффициентов")
			}
			
			input = strings.ReplaceAll(input, ",", ".")

			coeff[i][j], err = strconv.ParseFloat(input, 64)

			if err != nil {
				return 0, matrix{}, errors.New("коэффициенты должны быть числами")
			}
		}
	}

	return eps, matrix{size: size, coeff: coeff}, nil
}

func (reader FileReader) Read() (float64, matrix, error) {
	var path string
	fmt.Print("Введите название файла: ")

	fmt.Scanln(&path)

	return reader.ReadWithPath(path)

}

func (reader PreparedReader) Read() (float64, matrix, error) {
	return reader.ReadWithPath(reader.Path)
}

func (FileReader) ReadWithPath (path string) (float64, matrix, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeIrregular)

	if err != nil{
		return 0, matrix{}, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	line, _, err := reader.ReadLine()

	if err == io.EOF {
		return 0, matrix{}, errors.New("файл пустой")
	}

	if err != nil {
		return 0, matrix{}, err
	}

	size, err := strconv.Atoi(strings.Split(string(line), " ")[0])

	if err != nil {
		return 0, matrix{}, errors.New("размер матрицы - целое число")
	}

	if size <= 0 {
		return 0, matrix{}, errors.New("размер матрицы - положительное число")
	}

	eps, err := strconv.ParseFloat(strings.Split(string(line), " ")[1], 64)

	if err != nil {
		return 0, matrix{}, errors.New("точность - число")
	}

	if eps <= 0 {
		return 0, matrix{}, errors.New("точность - положительное число")
	}

	coeff := make([][]float64, size)

	for i:=0;i<size;i++{
		coeff[i] = make([]float64, size+1)
		line, _, err = reader.ReadLine()

		lineCoeff := strings.Split(string(line), " ")

		if err != nil{
			return 0, matrix{}, errors.New("недостаточно строк")
		}

		if len(lineCoeff) < size {
			return 0, matrix{}, errors.New("недостаточно коэффициентов в строке")
		}
		
		for j:=0;j<size+1;j++{

			value, err := strconv.ParseFloat(lineCoeff[j], 64)

			if err != nil {
				return 0, matrix{}, errors.New("коэффициенты должны быть числами")
			}

			coeff[i][j] = value
		}
	}

	return eps, matrix{size: size, coeff: coeff}, nil
}