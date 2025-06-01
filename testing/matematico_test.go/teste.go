// arquivo: matematica_test.go
package matematica

import "testing"

func TestSoma(t *testing.T) {
	resultado := Soma(2, 3)
	esperado := 5

	if resultado != esperado {
		t.Errorf("Esperado %d, mas recebeu %d", esperado, resultado)
	}
}
