package methods

import (
	"lab2/internal/functions"
	"lab2/internal/readers"
	"math"
)

func secantMethod(function functions.Function, readInfo readers.ReadInfo) (root float64, err error) {
	previousRoot := readInfo.Approx[0]
	root = readInfo.Approx[1]
	for math.Abs(previousRoot-root) > readInfo.Eps && math.Abs(function.Formula(root)) > readInfo.Eps {
		root = root - (root-previousRoot)*function.Formula(root)/(function.Formula(root)-function.Formula(previousRoot))
	}
	return
}
