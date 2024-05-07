package main

import "testing"

func TestComplexCalculation(t *testing.T) {
	//Criar uma tabela do caso de teste
	tests := []struct {
		name     string
		total    float64
		tax      float64
		expected float64
	}{
		{"When the total is greater than 100", 101.0, 0.1, 10.1},
		{"When the total is less than 100 and greater than 50", 51.0, 0.1, 51.1},
		{"When the total is less than 50", 40.0, 0.1, 40.1},
	}

	const floatTolerance = 0.0001

	for _, tt := range tests {
		t.Run(tt.name, func(te *testing.T) {
			result := complexCalculation(tt.total, tt.tax)
			if (result - tt.expected) > floatTolerance {
				te.Errorf("complexCalculation(%f, %f) = %f; want %f", tt.total, tt.tax, result, tt.expected)
			}
		})
	}
}
