package readers

import (
	"bufio"
	"errors"
	"fmt"
	"lab2/internal/functions"
	"lab2/internal/methods"
	"os"
	"strconv"
)

type ReadInfo struct {
	LeftBorder  float64
	RightBorder float64
	Eps         float64
	Approx      []float64 // может быть nil!
}

func ReadByConsole(function functions.Function, method methods.MethodInfo) (readInfo ReadInfo, err error) {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите левую границу промежутка: ")

	if !scanner.Scan() {
		err = errors.New("ожидался ввод")
		return
	}

	readInfo.LeftBorder, err = strconv.ParseFloat(scanner.Text(), 64)

	if err != nil {
		err = errors.New("ожидалось целое число")
		return
	}

	fmt.Print("Введите правую границу промежутка: ")

	if !scanner.Scan() {
		err = errors.New("ожидался ввод")
		return
	}

	readInfo.RightBorder, err = strconv.ParseFloat(scanner.Text(), 64)

	if err != nil {
		err = errors.New("ожидалось целое число")
		return
	}

	if readInfo.LeftBorder > readInfo.RightBorder {
		readInfo.LeftBorder, readInfo.RightBorder = readInfo.RightBorder, readInfo.LeftBorder
	}

	fmt.Print("Введите погрешность вычислений: ")

	if !scanner.Scan() {
		err = errors.New("ожидался ввод")
		return
	}

	readInfo.Eps, err = strconv.ParseFloat(scanner.Text(), 64)

	if err != nil {
		err = errors.New("ожидалось число")
		return
	}

	switch method.Id {
	case 1:
		readInfo.Approx = make([]float64, 2)
		fmt.Print("Введите начальное приближение: ")

		var isInputed bool

		if scanner.Scan() {
			readInfo.Approx[0], err = strconv.ParseFloat(scanner.Text(), 64)

			if err == nil {
				isInputed = true
			}
		}

		if !isInputed {
			fmt.Println("Было использовано значение по умолчанию")
			readInfo.Approx[0] = methods.GetFirstApprox(function, readInfo.LeftBorder, readInfo.RightBorder)
		}

		isInputed = false
		fmt.Print("Введите второе начальное приближение: ")

		if scanner.Scan() {
			readInfo.Approx[1], err = strconv.ParseFloat(scanner.Text(), 64)

			if err == nil {
				isInputed = true
			}
		}

		if !isInputed {
			fmt.Println("Было использовано значение по умолчанию")
			readInfo.Approx[1] = methods.GetSecondApprox(readInfo.Approx[0], readInfo.LeftBorder, readInfo.RightBorder)
		}

		err = nil
		
	}

	return
}
