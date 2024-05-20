package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestContainersTask(t *testing.T) {
	tests := []struct {
		inputFile      string
		expectedOutput string
	}{
		{"tests/test_file_1.txt", "tests/expected_output_1.txt"},
		{"tests/test_file_2.txt", "tests/expected_output_2.txt"},
		{"tests/test_file_3.txt", "tests/expected_output_3.txt"},
		{"tests/test_file_4.txt", "tests/expected_output_4.txt"},
		{"tests/test_file_5.txt", "tests/expected_output_5.txt"},
		{"tests/test_file_6.txt", "tests/expected_output_6.txt"},
		{"tests/test_file_7.txt", "tests/expected_output_7.txt"},
	}

	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			r, w, _ := os.Pipe()
			stdout := os.Stdout
			os.Stdout = w


			file, err := os.Open(tt.inputFile)
			if err!= nil {
				t.Fatalf("Failed to open input file: %v", err)
			}
			defer file.Close()

			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()
			os.Stdin = file

			main()
			
			w.Close()
			os.Stdout = stdout

			var buf bytes.Buffer
			buf.ReadFrom(r)
			actualOutput := buf.String()

			expectedOutput, err := os.ReadFile(tt.expectedOutput)
			if err != nil {
				t.Fatalf("Failed to read expected output file: %v", err)
			}

			if strings.TrimSpace(actualOutput) != strings.TrimSpace(string(expectedOutput)) {
				t.Errorf("Expected output:\n%s\n\nActual output:\n%s\n", expectedOutput, actualOutput)
			}
		})
	}
}