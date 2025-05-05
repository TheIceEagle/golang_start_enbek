// functions_test.go
package main

import (
	"fmt"
	"testing"
)

// TestSum проверяет функцию Sum
func TestSum(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{2, 3, 5},
		{-1, 1, 0},
		{0, 0, 0},
		{10, -5, 5},
		{-10, -5, -15},
	}
	
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Sum(%d,%d)", tt.a, tt.b), func(t *testing.T) {
			got := Sum(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Sum(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// TestAverage проверяет функцию Average
func TestAverage(t *testing.T) {
	tests := []struct {
		a, b, c, want float64
	}{
		{1, 2, 3, 2},
		{0, 0, 0, 0},
		{-1, 0, 1, 0},
		{5.5, 10.5, 5, 7},
		{-10, 10, 0, 0},
	}
	
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Average(%.1f,%.1f,%.1f)", tt.a, tt.b, tt.c), func(t *testing.T) {
			got := Average(tt.a, tt.b, tt.c)
			if got != tt.want {
				t.Errorf("Average(%.1f, %.1f, %.1f) = %.1f; want %.1f", 
					tt.a, tt.b, tt.c, got, tt.want)
			}
		})
	}
}

// TestGreeting проверяет функцию Greeting
func TestGreeting(t *testing.T) {
	tests := []struct {
		name, want string
	}{
		{"Иван", "Привет, Иван!"},
		{"Мария", "Привет, Мария!"},
		{"", "Привет, !"},
		{"Алексей", "Привет, Алексей!"},
	}
	
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Greeting(%s)", tt.name), func(t *testing.T) {
			got := Greeting(tt.name)
			if got != tt.want {
				t.Errorf("Greeting(%q) = %q; want %q", tt.name, got, tt.want)
			}
		})
	}
}

// TestCalculateRoomPrice проверяет функцию CalculateRoomPrice
func TestCalculateRoomPrice(t *testing.T) {
	tests := []struct {
		rate   float64
		nights uint
		want   float64
	}{
		{100, 3, 300},
		{50.5, 2, 101},
		{0, 5, 0},
		{200, 0, 0},
		{123.45, 7, 864.15},
	}
	
	for _, tt := range tests {
		t.Run(fmt.Sprintf("CalculateRoomPrice(%.1f,%d)", tt.rate, tt.nights), func(t *testing.T) {
			got := CalculateRoomPrice(tt.rate, tt.nights)
			if got != tt.want {
				t.Errorf("CalculateRoomPrice(%.2f, %d) = %.2f; want %.2f", 
					tt.rate, tt.nights, got, tt.want)
			}
		})
	}
}

// TestApplyDiscount проверяет функцию ApplyDiscount
func TestApplyDiscount(t *testing.T) {
	tests := []struct {
		price    float64
		discount float64
		want     float64
	}{
		{100, 10, 90},
		{200, 0, 200},
		{50, 100, 0},
		{80, 25, 60},
		{1000, 33.3, 667},
	}
	
	for _, tt := range tests {
		t.Run(fmt.Sprintf("ApplyDiscount(%.1f,%.1f)", tt.price, tt.discount), func(t *testing.T) {
			got := ApplyDiscount(tt.price, tt.discount)
			// Округляем до целых для избежания проблем с плавающей точкой
			if int(got) != int(tt.want) {
				t.Errorf("ApplyDiscount(%.2f, %.2f) = %.2f; want %.2f", 
					tt.price, tt.discount, got, tt.want)
			}
		})
	}
}

// TestFormatRoomInfo проверяет функцию FormatRoomInfo
func TestFormatRoomInfo(t *testing.T) {
	tests := []struct {
		roomNumber int
		capacity   int
		nights     int
		want       string
	}{
		{101, 2, 3, "Номер 101: вместимость 2 человека, 3 ночей"},
		{205, 1, 1, "Номер 205: вместимость 1 человек, 1 ночь"},
		{310, 4, 0, "Номер 310: вместимость 4 человека, 0 ночей"},
		{402, 5, 7, "Номер 402: вместимость 5 человек, 7 ночей"},
	}
	
	for _, tt := range tests {
		t.Run(fmt.Sprintf("FormatRoomInfo(%d,%d,%d)", tt.roomNumber, tt.capacity, tt.nights), func(t *testing.T) {
			got := FormatRoomInfo(tt.roomNumber, tt.capacity, tt.nights)
			if got != tt.want {
				t.Errorf("FormatRoomInfo(%d, %d, %d) = %q; want %q", 
					tt.roomNumber, tt.capacity, tt.nights, got, tt.want)
			}
		})
	}
}

// TestCalculateOccupancyRate проверяет функцию CalculateOccupancyRate
func TestCalculateOccupancyRate(t *testing.T) {
	tests := []struct {
		occupied int
		total    int
		want     float64
	}{
		{50, 100, 50},
		{0, 100, 0},
		{100, 100, 100},
		{25, 50, 50},
		{75, 150, 50},
	}
	
	for _, tt := range tests {
		t.Run(fmt.Sprintf("CalculateOccupancyRate(%d,%d)", tt.occupied, tt.total), func(t *testing.T) {
			got := CalculateOccupancyRate(tt.occupied, tt.total)
			if got != tt.want {
				t.Errorf("CalculateOccupancyRate(%d, %d) = %.1f; want %.1f", 
					tt.occupied, tt.total, got, tt.want)
			}
		})
	}
}

// TestDetermineOccupancyLevel проверяет функцию DetermineOccupancyLevel
func TestDetermineOccupancyLevel(t *testing.T) {
	tests := []struct {
		rate float64
		want string
	}{
		{20, "Низкий"},
		{30, "Средний"},
		{45, "Средний"},
		{60, "Средний"},
		{75, "Высокий"},
		{0, "Низкий"},
		{100, "Высокий"},
	}
	
	for _, tt := range tests {
		t.Run(fmt.Sprintf("DetermineOccupancyLevel(%.1f)", tt.rate), func(t *testing.T) {
			got := DetermineOccupancyLevel(tt.rate)
			if got != tt.want {
				t.Errorf("DetermineOccupancyLevel(%.1f) = %q; want %q", 
					tt.rate, got, tt.want)
			}
		})
	}
}

// TestGetRecommendation проверяет функцию GetRecommendation
func TestGetRecommendation(t *testing.T) {
	tests := []struct {
		level string
		want  string
	}{
		{"Низкий", "Запустите программу скидок"},
		{"Средний", "Обычный режим работы"},
		{"Высокий", "Повысьте цены"},
	}
	
	for _, tt := range tests {
		t.Run(fmt.Sprintf("GetRecommendation(%s)", tt.level), func(t *testing.T) {
			got := GetRecommendation(tt.level)
			if got != tt.want {
				t.Errorf("GetRecommendation(%q) = %q; want %q", 
					tt.level, got, tt.want)
			}
		})
	}
}

// TestFindMax проверяет функцию FindMax
func TestFindMax(t *testing.T) {
	tests := []struct {
		numbers []int
		want    int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{5, 3, 8, 2}, 8},
		{[]int{-1, -5, -3}, -1},
		{[]int{0}, 0},
		{[]int{}, 0}, // Проверка на пустой слайс
		{[]int{42}, 42},
		{[]int{7, 7, 7}, 7}, // Все элементы равны
	}
	
	for i, tt := range tests {
		testName := fmt.Sprintf("FindMax case %d", i)
		t.Run(testName, func(t *testing.T) {
			got := FindMax(tt.numbers...)
			if got != tt.want {
				t.Errorf("FindMax(%v) = %d; want %d", tt.numbers, got, tt.want)
			}
		})
	}
}

// TestConcatenate проверяет функцию Concatenate
func TestConcatenate(t *testing.T) {
	tests := []struct {
		separator string
		parts     []string
		want      string
	}{
		{", ", []string{"Иван", "Мария", "Алексей"}, "Иван, Мария, Алексей"},
		{"-", []string{"a", "b", "c"}, "a-b-c"},
		{"", []string{"x", "y", "z"}, "xyz"},
		{":", []string{"Привет"}, "Привет"},
		{" и ", []string{"яблоки", "апельсины"}, "яблоки и апельсины"},
		{" ", []string{}, ""}, // Проверка пустого слайса
	}
	
	for i, tt := range tests {
		testName := fmt.Sprintf("Concatenate case %d", i)
		t.Run(testName, func(t *testing.T) {
			got := Concatenate(tt.separator, tt.parts...)
			if got != tt.want {
				t.Errorf("Concatenate(%q, %v) = %q; want %q", 
					tt.separator, tt.parts, got, tt.want)
			}
		})
	}
}