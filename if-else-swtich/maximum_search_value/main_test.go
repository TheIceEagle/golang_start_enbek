package main

import (
	"testing"
)

func TestFindMax(t *testing.T) {

	tests := []struct {
		name     string
		a, b, c  int
		expected int
	}{
		{"Максимум - третье число", 10, 20, 30, 30},
		{"Максимум - первое число", 50, 10, 20, 50},
		{"Максимум - второе число", 15, 45, 5, 45},
		{"Два максимальных числа", 100, 100, 1, 100},
		{"Все числа равны", 7, 7, 7, 7},
		{"Отрицательные числа", -5, -2, -10, -2},
		{"Разные знаки", -10, 0, 5, 5},
	}


	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FindMax(test.a, test.b, test.c)
			if result != test.expected {
				t.Errorf("Для чисел %d, %d, %d получено максимальное значение %d, ожидалось %d", 
					test.a, test.b, test.c, result, test.expected)
			}
		})
	}
}