package methods

import "lab2/internal/functions"

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
	Action func(functions.Function) (root float64, err error)
}

func HasIntervalRoot(function functions.Function, a int, b int) bool {
	return HasIntervalRoots(function, a, b) && function.Derivative(float64(a))*function.Derivative(float64(b)) > 0
}

func HasIntervalRoots(function functions.Function, a int, b int) bool {
	return a*b < 0
}

func GetFisrtApprox(function functions.Function, a float64, b float64) (float64) {
	if function.Formula(a)*function.Derivative2(a) > 0 {
		return a
	} else {
		return b
	}
}

func GetSecondApprox(firstApprox float64, a float64, b float64) (float64) {
	return firstApprox+(b-a)/10
}