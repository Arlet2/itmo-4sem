package methods

import (
	"lab2/internal/functions"
	"lab2/internal/readers"
)

var Methods = []MethodInfo{
	{
		Id:     0,
		Name:   "Метод половинного деления",
		Action: halfDivisionMethod,
	},
	{
		Id:     1,
		Name:   "Метод секущих",
		Action: secantMethod,
	},
	{
		Id:     2,
		Name:   "Метод простых итераций",
		Action: simpleIterationsMethod,
	},
}

type MethodInfo struct {
	Id     int
	Name   string
	Action func(functions.Function, readers.ReadInfo) (root float64, err error)
}