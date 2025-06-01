package matrix

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	a, _ := NewMatrix([][]float64{
		{1, 2},
		{3, 4},
	})
	b, _ := NewMatrix([][]float64{
		{2, 0},
		{1, 2},
	})
	expected := [][]float64{
		{4, 4},
		{10, 8},
	}
	res, err := a.Multiply(b)
	if err != nil {
		t.Fatalf("Erro na multiplicação: %v", err)
	}
	for i := range expected {
		for j := range expected[i] {
			if res.Data[i][j] != expected[i][j] {
				t.Errorf("Valor esperado %.2f, obtido %.2f", expected[i][j], res.Data[i][j])
			}
		}
	}
}

func BenchmarkMultiply(b *testing.B) {
	m1, _ := NewMatrix(generateMatrix(200, 300))
	m2, _ := NewMatrix(generateMatrix(300, 400))
	for i := 0; i < b.N; i++ {
		_, _ = m1.Multiply(m2)
	}
}

func generateMatrix(rows, cols int) [][]float64 {
	mat := make([][]float64, rows)
	for i := range mat {
		mat[i] = make([]float64, cols)
		for j := range mat[i] {
			mat[i][j] = float64(i + j)
		}
	}
	return mat
}
