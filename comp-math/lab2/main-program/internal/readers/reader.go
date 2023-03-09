package readers

import (
	"bufio"
	"errors"
	"fmt"
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

func ReadByConsole(method methods.MethodInfo) (readInfo ReadInfo, err error) {

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
		// todo: добавить значения по умолчанию
		fmt.Print("Введите начальное приближение: ")

		if !scanner.Scan() {
			err = errors.New("ожидался ввод")
			return
		}

		readInfo.Approx[0], err = strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			err = errors.New("ожидалось число")
			return
		}

		fmt.Print("Введите второе начальное приближение: ")

		if !scanner.Scan() {
			err = errors.New("ожидался ввод")
			return
		}

		readInfo.Approx[1], err = strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			err = errors.New("ожидалось число")
			return
		}
		
	}

	return
}