package main

import (
	"testing"
)

func TestGetTemperatureMessage(t *testing.T) {
	tests := []struct {
		name     string
		temp     int
		expected string
	}{
		{"Отрицательная температура", -10, "Очень холодно"},
		{"Граничное значение - минус один", -1, "Очень холодно"},
		{"Граничное значение - ноль", 0, "Прохладно"},
		{"Прохладная температура", 7, "Прохладно"},
		{"Граничное значение - пятнадцать", 15, "Прохладно"},
		{"Граничное значение - шестнадцать", 16, "Тепло"},
		{"Тёплая температура", 25, "Тепло"},
		{"Граничное значение - тридцать", 30, "Тепло"},
		{"Граничное значение - тридцать один", 31, "Жарко"},
		{"Жаркая температура", 40, "Жарко"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetTemperatureMessage(test.temp)
			if result != test.expected {
				t.Errorf("для температуры %d °C ожидалось '%s', получено '%s'", 
					test.temp, test.expected, result)
			}
		})
	}
}


