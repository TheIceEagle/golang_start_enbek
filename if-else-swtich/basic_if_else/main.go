package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	temperature := rand.Intn(41)
	fmt.Println("Температура:", temperature, "°C -", GetTemperatureMessage(temperature))
}


// TODO: Реализуйте эту функцию согласно заданию
func GetTemperatureMessage(temp int) string {
	// Удалите эту строку и напишите своё решение
	return "Напишите свою реализацию здесь"
}