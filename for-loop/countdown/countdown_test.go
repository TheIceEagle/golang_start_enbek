package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCountdown(t *testing.T) {

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w


	go func() {
		Countdown()
		w.Close() 
	}()

	
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = old 

	output := buf.String()

	expectedOutput := "10\n9\n8\n7\n6\n5\n4\n3\n2\n1\nПуск!\n"
	if output != expectedOutput {
		t.Errorf("Ожидался вывод:\n%s\nПолучен вывод:\n%s", expectedOutput, output)
	}

	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != 11 {
		t.Errorf("Ожидалось 11 строк (10 чисел и 'Пуск!'), получено %d", len(lines))
	}

	for i := 0; i < 10; i++ {
		expected := fmt.Sprintf("%d", 10-i)
		if lines[i] != expected {
			t.Errorf("Строка %d: ожидалось '%s', получено '%s'", i+1, expected, lines[i])
		}
	}

	if lines[10] != "Пуск!" {
		t.Errorf("Последняя строка: ожидалось 'Пуск!', получено '%s'", lines[10])
	}
}
