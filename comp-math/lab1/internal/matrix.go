package internal

import "fmt"

type matrix struct {
	size int
	coeff [][]float64
}

func (m matrix) Print() {
	fmt.Printf("Matrix %dx%d:\n", m.size, m.size)

	for i:=0;i<m.size;i++{
		for j:=0;j<m.size+1;j++{
			fmt.Print(m.coeff[i][j], " ")
		}
		fmt.Println()
	}
}

func (m *matrix) MulRow(rowIndex int, value float64) {
	for index := range m.coeff[rowIndex] {
		m.coeff[rowIndex][index] *= value
	}
}