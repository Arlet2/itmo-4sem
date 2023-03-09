package main

import (
	"bufio"
	"fmt"
	"lab2/internal/functions"
	"lab2/internal/methods"
	"lab2/internal/readers"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	//todo: добавить разделение на систему уравнений и одно уравнение
	fmt.Println("Выберите функцию из предложенных:")
	for index, function := range functions.Functions {
		fmt.Printf("%d: f(x) = "+function.Text+"\n", index+1)
	}

	if !scanner.Scan() {
		fmt.Println("Ожидался ввод...")
		return
	}

	input, err := strconv.ParseInt(scanner.Text(), 10, 64)

	if err != nil {
		fmt.Println("Ожидалось целое число...")
		return
	}

	if int(input-1) < 0 || int(input) > len(functions.Functions) {
		fmt.Printf("Ожидалось число в промежутке от 1 до %d\n", len(functions.Functions))
		return
	}

	function := functions.Functions[int(input-1)]

	fmt.Println("Вы выбрали функцию f(x) = " + function.Text)

	fmt.Println("Выберите метод:")

	for index, method := range methods.Methods {
		fmt.Printf("%d) "+method.Name+"\n", index+1)
	}

	if !scanner.Scan() {
		fmt.Println("Ожидался ввод...")
		return
	}

	input, err = strconv.ParseInt(scanner.Text(), 10, 64)

	if err != nil {
		fmt.Println("Ожидалось целое число...")
		return
	}

	if int(input-1) < 0 || int(input-1) >= len(methods.Methods) {
		fmt.Printf("Ожидалось число в промежутке от 1 до %d\n", len(methods.Methods))
		return
	}

	method := methods.Methods[input-1]
	fmt.Println("Вы выбрали " + method.Name)

	fmt.Println("Выберите: 1) ввод из файла; 2) ввод из консоли; 3) ввод из lab2.txt")

	if !scanner.Scan() {
		fmt.Println("Ожидался ввод...")
		return
	}

	input, err = strconv.ParseInt(scanner.Text(), 10, 64)

	if err != nil {
		fmt.Println("Ожидалось целое число")
		return
	}

	var readInfo readers.ReadInfo

	if input == 1 {
		fmt.Print("Введите путь до файла: ")
		if !scanner.Scan() {
			fmt.Println("Ожидался ввод...")
			return
		}

		readInfo, err = readers.ReadByFile(scanner.Text(), function, method.Id)
	} else if input == 2 {
		readInfo, err = readers.ReadByConsole(function, method.Id)
	} else if input == 3 {
		readInfo, err = readers.ReadByFile("lab3.txt", function, method.Id)
	} else {
		fmt.Println("Ожидалось число от 1 до 3")
		return
	}

	if err != nil {
		fmt.Println("Ошибка: " + err.Error())
		return
	}

	fmt.Println(readInfo)

	if !functions.HasIntervalRoots(function, readInfo.LeftBorder, readInfo.RightBorder) {
		fmt.Println("ВНИМАНИЕ! На данном интервале ВОЗМОЖНО нет корней. Результаты могут быть некорректны")
	} else if !functions.HasIntervalRoot(function, readInfo.LeftBorder, readInfo.RightBorder) {
		fmt.Println("ВНИМАНИЕ! На данном интервале НЕСКОЛЬКО корней. Будет найден произвольный корень")
	}

	root, err := method.Action(function, readInfo)

	if err != nil {
		fmt.Println("Ошибка: " + err.Error())
		return
	}

	fmt.Printf("Найденный корень равен: %f\n", root)

}
