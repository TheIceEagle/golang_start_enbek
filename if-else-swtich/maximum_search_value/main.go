package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	a := rand.Intn(100) + 1
	b := rand.Intn(100) + 1
	c := rand.Intn(100) + 1
	
	fmt.Printf("Числа: %d, %d, %d\n", a, b, c)
	
	max := FindMax(a, b, c)
	
	fmt.Printf("Максимальное значение: %d\n", max)
}

// FindMax возвращает максимальное значение из трёх чисел
// TODO: Реализуйте эту функцию согласно заданию
func FindMax(a, b, c int) int {
	// Удалите эту строку и напишите своё решение
	return 0 
}