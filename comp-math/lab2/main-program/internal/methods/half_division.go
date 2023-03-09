package methods

import (
	"lab2/internal/functions"
	"lab2/internal/readers"
	"math"
)

func halfDivisionMethod(function functions.Function, readInfo readers.ReadInfo) (root float64, err error) {
	root = (readInfo.RightBorder + readInfo.LeftBorder) / 2
	for math.Abs(readInfo.RightBorder-readInfo.LeftBorder) > readInfo.Eps && math.Abs(function.Formula(root)) > readInfo.Eps {
		if function.Formula(readInfo.LeftBorder)*function.Formula(root) < 0 {
			readInfo.RightBorder = root
		} else if function.Formula(readInfo.RightBorder)*function.Formula(root) < 0 {
			readInfo.LeftBorder = root
		} else {
			readInfo.LeftBorder += (readInfo.RightBorder + readInfo.LeftBorder) / 100
			readInfo.RightBorder -= (readInfo.RightBorder + readInfo.LeftBorder) / 100
		}

		root = (readInfo.RightBorder + readInfo.LeftBorder) / 2
	}
	return
}
