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

func HasIntervalRoot(function functions.Function, a int, b int) bool {
	return HasIntervalRoots(function, a, b) && function.Derivative(float64(a))*function.Derivative(float64(b)) > 0
}

func HasIntervalRoots(function functions.Function, a int, b int) bool {
	return a*b < 0
}
