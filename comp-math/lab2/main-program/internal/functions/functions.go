package functions

import "math"

var Functions = []Function{
	{
		Text:       "2022x^2 - 2023x + 1",
		Formula:    func(x float64) float64 { return 2022*math.Pow(x, 2) - 2023*x + 1 },
		Derivative: func(x float64) float64 { return 8088*x - 2023 },
	},
	{
		Text:       "pi - x",
		Formula:    func(x float64) float64 { return math.Pi - x },
		Derivative: func(x float64) float64 { return -1 }},
	{
		Text:       "x - 20",
		Formula:    func(x float64) float64 { return x - 20 },
		Derivative: func(x float64) float64 { return 1 },
	},
	{
		Text:       "ln x + 20",
		Formula:    func(x float64) float64 { return math.Log(x) + 20 },
		Derivative: func(x float64) float64 { return 1 / x },
	},
	{
		Text:       "sin 2x",
		Formula:    func(x float64) float64 { return math.Sin(2 * x) },
		Derivative: func(x float64) float64 { return 2 * math.Cos(2*x) },
	},
	{
		Text:       "x^3 - 4x^2 - 7x + 10",
		Formula:    func(x float64) float64 { return math.Pow(x, 3) - 4*math.Pow(x, 2) - 7*x + 10 },
		Derivative: func(x float64) float64 { return 3*math.Pow(x, 2) - 8*x - 7 },
	},
}

type Function struct {
	Text       string
	Formula    func(float64) float64
	Derivative func(float64) float64
}