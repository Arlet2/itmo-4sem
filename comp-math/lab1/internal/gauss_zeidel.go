package internal

import (
	"fmt"
	"math"
)

/*
	1. Ищем максимальный элемент в строке, записываем индекс
	2. Проверяем по условию в каждой строке (сумма остальных элементов меньше максимального)
	3. Добавляем максимальный индекс каждого элемента в коллекцию
	4. Проверяем: если индексы различны, то возможно перестроить матрицу так, чтобы она обладала диагональным преобладанием, иначе нет
*/
func (m *matrix) TryToCreateDiagonalDominance() bool {

	var maxCoeff float64
	var maxIndex int
	var rowSum float64
	oneStrictExist := false

	maxIndexes := map[int][]float64{}

	for i:=0; i<m.size; i++ {
		maxCoeff = math.Abs(m.coeff[i][0])
		maxIndex = 0
		// поиск максимального элемента
		for j := 0; j<m.size;j++ {
			if maxCoeff < math.Abs(m.coeff[i][j]) {
				maxIndex = j
				maxCoeff = math.Abs(m.coeff[i][j])
			}
		}
		rowSum = 0
		// проверка условия
		for j := 0; j<m.size; j++ {
			// игнорируем значение максимального элемента
			if j == maxIndex {
				continue
			}
			rowSum += m.coeff[i][j]
		}

		if maxCoeff < rowSum {
			return false
		}

		if maxCoeff > rowSum {
			oneStrictExist = true
		}

		_, exist := maxIndexes[maxIndex]

		if exist {
			return false
		}

		maxIndexes[maxIndex] = m.coeff[i]
	}

	for key, val := range maxIndexes {
		fmt.Printf("[%d] = %f\n", key, val)
	}

	if len(maxIndexes) != m.size {
		return false
	}

	if !oneStrictExist {
		fmt.Println("Ни для одного уравнения не выполняется строгое преобладание диагонали.")
		return false
	}

	for i:=0;i<m.size;i++ {
		m.coeff[i] = maxIndexes[i]
	}

	return true
}

