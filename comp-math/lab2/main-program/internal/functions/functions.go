package functions

import "math"

var Functions = []Function{
	{
		Text:        "2022x^2 - 2023x + 1",
		Formula:     func(x float64) float64 { return 2022*math.Pow(x, 2) - 2023*x + 1 },
		Derivative:  func(x float64) float64 { return 4044*x - 2023 },
		Derivative2: func(x float64) float64 { return 4044 },
	},
	{
		Text:        "pi - x",
		Formula:     func(x float64) float64 { return math.Pi - x },
		Derivative:  func(x float64) float64 { return -1 },
		Derivative2: func(x float64) float64 { return 0 },
	},
	{
		Text:        "x - 20",
		Formula:     func(x float64) float64 { return x - 20 },
		Derivative:  func(x float64) float64 { return 1 },
		Derivative2: func(x float64) float64 { return 0 },
	},
	{
		Text:        "100*ln x + 20",
		Formula:     func(x float64) float64 { return 100*math.Log(x) + 20 },
		Derivative:  func(x float64) float64 { return 100 / x },
		Derivative2: func(x float64) float64 { return -100 / math.Pow(x, 2) },
	},
	{
		Text:        "sin 2x",
		Formula:     func(x float64) float64 { return math.Sin(2 * x) },
		Derivative:  func(x float64) float64 { return 2 * math.Cos(2*x) },
		Derivative2: func(x float64) float64 { return -4 * math.Sin(2*x) },
	},
	{
		Text:        "x^3 - 4x^2 - 7x + 10",
		Formula:     func(x float64) float64 { return math.Pow(x, 3) - 4*math.Pow(x, 2) - 7*x + 10 },
		Derivative:  func(x float64) float64 { return 3*math.Pow(x, 2) - 8*x - 7 },
		Derivative2: func(x float64) float64 { return 6*x - 8 },
	},
}

type Function struct {
	Text        string
	Formula     func(float64) float64
	Derivative  func(float64) float64
	Derivative2 func(float64) float64
}

func HasIntervalRoot(function Function, a float64, b float64) bool {
	return HasIntervalRoots(function, a, b) && isMonotone(function, a, b)
}

func HasIntervalRoots(function Function, a float64, b float64) bool {
	return function.Formula(a)*function.Formula(b) <= 0
}

func isMonotone(function Function, a float64, b float64) bool {
	for i := a; i <= b; i += (b - a) / 100 {
		if function.Derivative(a)*function.Derivative(i) < 0 {
			return false
		}
	}

	return true
}

func GetFirstApprox(function Function, a float64, b float64) float64 {
	if function.Formula(a)*function.Derivative2(a) > 0 {
		return a
	} else {
		return b
	}
}

func GetSecondApprox(firstApprox float64, a float64, b float64) float64 {
	if firstApprox == a {
		return firstApprox + (b-a)/10
	} else {
		return firstApprox - (b-a)/10
	}

}
