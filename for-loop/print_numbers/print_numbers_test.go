package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestPrintNumbers(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintNumbers()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	expectedOutput := "1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n"
	if output != expectedOutput {
		t.Errorf("Ожидался вывод:\n%s\nПолучен вывод:\n%s", expectedOutput, output)
	}

	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != 10 {
		t.Errorf("Ожидалось 10 строк вывода, получено %d", len(lines))
	}

	for i, line := range lines {
		expectedLine := i + 1
		if line != string(rune('0'+expectedLine)) {
			t.Errorf("Строка %d: ожидалось '%d', получено '%s'", 
				i+1, expectedLine, line)
		}
	}
}