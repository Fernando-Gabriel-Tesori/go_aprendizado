package matrix

import (
	"errors"
	"fmt"
	"sync"
)

type Matrix struct {
	Data [][]float64
	Rows int
	Cols int
}

// NewMatrix cria uma matriz genérica
func NewMatrix(data [][]float64) (*Matrix, error) {
	if len(data) == 0 || len(data[0]) == 0 {
		return nil, errors.New("matriz vazia")
	}
	rows := len(data)
	cols := len(data[0])
	for _, row := range data {
		if len(row) != cols {
			return nil, fmt.Errorf("linhas com tamanhos diferentes")
		}
	}
	return &Matrix{Data: data, Rows: rows, Cols: cols}, nil
}

// Multiply faz multiplicação concorrente de matrizes
func (m *Matrix) Multiply(other *Matrix) (*Matrix, error) {
	if m.Cols != other.Rows {
		return nil, errors.New("dimensões incompatíveis para multiplicação")
	}

	result := make([][]float64, m.Rows)
	for i := range result {
		result[i] = make([]float64, other.Cols)
	}

	var wg sync.WaitGroup
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < other.Cols; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				sum := 0.0
				for k := 0; k < m.Cols; k++ {
					sum += m.Data[i][k] * other.Data[k][j]
				}
				result[i][j] = sum
			}(i, j)
		}
	}
	wg.Wait()
	return NewMatrix(result)
}
