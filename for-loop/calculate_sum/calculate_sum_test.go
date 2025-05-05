package main

import (
	"testing"
)

func TestCalculateSum(t *testing.T) {

	result := CalculateSum()

	
	expected := 5050 
	if result != expected {
		t.Errorf("Ожидалась сумма %d, получено %d", expected, result)
	}
}


func TestCalculateSumBoundaries(t *testing.T) {
	
	sum1to10 := func() int {
		sum := 0
		num := 1
		for num <= 10 {
			sum += num
			num++
		}
		return sum
	}()
	
	if sum1to10 != 55 {
		t.Errorf("Ожидалась сумма чисел от 1 до 10: 55, получено %d", sum1to10)
	}
	
	
	sum1to1000 := func() int {
		sum := 0
		num := 1
		for num <= 1000 {
			sum += num
			num++
		}
		return sum
	}()
	
	if sum1to1000 != 500500 {
		t.Errorf("Ожидалась сумма чисел от 1 до 1000: 500500, получено %d", sum1to1000)
	}
}