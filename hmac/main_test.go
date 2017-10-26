package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSign(t *testing.T) {
	//TODO: write unit test cases for sign()
	//use strings.NewReader() to get an io.Reader
	//interface over a simple string
	//https://golang.org/pkg/strings/#NewReader
	cases := []struct {
		name           string
		input1         string
		input2         io.Reader
		expectedOutput string
	}{
		{
			name:           "empty string",
			input1:         "password",
			input2:         strings.NewReader("Hello"),
			expectedOutput: "",
		},
	}

	for _, c := range cases {
		output, err := sign(c.input1, c.input2)
		if err != nil {
			fmt.Printf("error signing: %v", err)
			os.Exit(exitCodeProcessing)
		}
		if output != c.expectedOutput {
			t.Errorf("%s: got %s but expected %s", c.name, output, c.expectedOutput)
		}
	}
}

func TestVerify(t *testing.T) {
	//TODO: write until test cases for verify()
	//use strings.NewReader() to get an io.Reader
	//interface over a simple string
	//https://golang.org/pkg/strings/#NewReader
}
