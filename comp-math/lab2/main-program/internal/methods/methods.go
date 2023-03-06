package methods

import "lab2/internal/functions"

var Methods = []MethodInfo{
	{
		Name:   "Метод половинного деления",
		Action: halfDivisionMethod,
	},
	{
		Name:   "Метод секущих",
		Action: secantMethod,
	},
	{
		Name:   "Метод простых итераций",
		Action: simpleIterationsMethod,
	},
}

type MethodInfo struct {
	Name   string
	Action func(functions.Function) (root float64, err error)
}
