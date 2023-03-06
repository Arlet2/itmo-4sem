package main

import (
	"bufio"
	"fmt"
	"lab2/internal/functions"
	"lab2/internal/methods"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
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

	if int(input-1) < 0 || int(input) >= len(functions.Functions) {
		fmt.Printf("Ожидалось число в промежутке от 1 до %d\n", len(functions.Functions))
		return
	}

	fmt.Println("Вы выбрали функцию f(x) = " + functions.Functions[int(input-1)].Text)

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

	if int(input-1) < 0 || int(input) >= len(methods.Methods) {
		fmt.Printf("Ожидалось число в промежутке от 1 до %d\n", len(methods.Methods))
		return
	}

	fmt.Println("Вы выбрали "+ methods.Methods[input-1].Name)
}
