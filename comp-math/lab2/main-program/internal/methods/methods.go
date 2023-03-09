package methods

import (
	"lab2/internal/functions"
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
	Action func(functions.Function, float64, float64) (root float64, err error)
}

func HasIntervalRoot(function functions.Function, a float64, b float64) bool {
	return HasIntervalRoots(function, a, b) && isMonotone(function, a, b)
}

func HasIntervalRoots(function functions.Function, a float64, b float64) bool {
	return function.Formula(a)*function.Formula(b) < 0
}

func isMonotone(function functions.Function, a float64, b float64) bool {
	for i := a; i <= b; i+= (b-a)/100 {
		if function.Derivative(a)*function.Derivative(i) < 0 {
			return false
		}
	}

	return true
}

func GetFirstApprox(function functions.Function, a float64, b float64) (float64) {
	if function.Formula(a)*function.Derivative2(a) > 0 {
		return a
	} else {
		return b
	}
}

func GetSecondApprox(firstApprox float64, a float64, b float64) (float64) {
	if firstApprox == a {
		return firstApprox+(b-a)/10
	} else {
		return firstApprox-(b-a)/10
	}
	
}