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
func (m matrix) TryToCreateDiagonalDominance() bool {

	var maxCoeff float64
	var maxIndex int
	var rowSum float64
	oneStrictExist := false

	maxIndexes := map[int][]float64{}

	for i := 0; i < m.size; i++ {
		maxCoeff = math.Abs(m.coeff[i][0])
		maxIndex = 0
		// поиск максимального элемента
		for j := 0; j < m.size; j++ {
			if maxCoeff < math.Abs(m.coeff[i][j]) {
				maxIndex = j
				maxCoeff = math.Abs(m.coeff[i][j])
			}
		}
		rowSum = 0
		// проверка условия
		for j := 0; j < m.size; j++ {
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

	if len(maxIndexes) != m.size {
		return false
	}

	if !oneStrictExist {
		fmt.Println("Ни для одного уравнения не выполняется строгое преобладание диагонали.")
		return false
	}

	for i := 0; i < m.size; i++ {
		m.coeff[i] = maxIndexes[i]
	}

	return true
}

func (m matrix) IsNormaCorrect() (bool) {
	fmt.Println("Вычисление нормы для матрицы: ")
	m.Print()
	var sum float64
	for i := 0; i < 0; i++ {
		sum = 0
		for _, value := range m.coeff[i] {
			sum += math.Abs(value)
		}
		if (sum >= 1) {
			return false
		}
	}
	return true
}

func (m matrix) UseGaussZeidel(eps float64, withTrace bool) {

	var specialMatrix matrix
	coeff := make([][]float64, m.size)
	// приводим матрицу к виду x = Cx + d
	for index := range m.coeff {
		m.MulRow(index, 1/m.coeff[index][index])

		for j := 0; j < m.size; j++ {
			if index == j {
				m.coeff[index][j] = 0
				continue
			}
			m.coeff[index][j] *= -1
		}
		coeff[index] = []float64{}

		for jndex, value := range m.coeff[index] {
			if jndex == index {
				continue
			}
			coeff[index] = append(coeff[index], value)
		}
	}

	specialMatrix = matrix{size: m.size, coeff: coeff}

	fmt.Println("\nМатрица вида x = Cx+d:")
	specialMatrix.Print()

	if m.IsNormaCorrect() {
		fmt.Println("Норма матрицы < 1. Условия сходимости выполняются")
	} else {
		fmt.Println("Норма матрицы не выполняет условия сходимости")
	}

	counter := 0
	diff := eps+1
	var last float64
	var shift int

	d := make([]float64, m.size)
	faults := make([]float64, m.size)

	// задаём начальное приближение
	for index := range specialMatrix.coeff {
		d[index] = specialMatrix.coeff[index][specialMatrix.size-1]
	}

	if (withTrace) {
		fmt.Println("Итерации:")
	}

	for diff >= eps {
		if (withTrace) {
			fmt.Println(d, "diff:",diff)
		}
		counter++
		diff = 0

		for i := 0; i < specialMatrix.size; i++ {
			last = d[i]
			d[i] = 0
			shift = 0
			for j := 0; j < m.size; j++ {
				if j == i {
					shift = 1
				}

				if j+1 == m.size {
					d[i] += specialMatrix.coeff[i][j]
					continue
				}

				d[i] += specialMatrix.coeff[i][j] * d[j+shift]
			}
			//d[i] = float64(float32(d[i]))
			faults[i] = math.Abs(d[i]-last)
			diff = math.Max(diff, faults[i])

		}

	}
	fmt.Println("\nВыполнено за", counter, "итерации(-й).")
	fmt.Println("Вектор неизвестных:", d)
	fmt.Println("Вектор погрешностей:", faults)

}
