package main

import "testing"
import "fmt"

func Test(t *testing.T) {
	var tests = []struct {
		input_path    string
		expected_path string
	}{
		{"D/go/code/../src/warcluster/tests/first/../../", "/D/go/src/warcluster/"},
		{"/python/movies/episode1/../../lectures/lecture1/examples/../code/../../../mostImportant/MonthyPython/quotes/..", "/python/mostImportant/MonthyPython/"},
		{"/", "/"},
		{"./..", "/"},
		{"..", "/"},
		{"../", "/"},
		{"/..", "/"},
		{"", "/"},
		{"/.", "/"},
		{".", "/"},
		{"./", "/"},
		{"/.", "/"},
		{"//", "/"},
		{"lala", "/lala/"},
		{"//path/to/your/heart", "/path/to/your/heart/"},
		{"path/tome/nowhere//", "/path/tome/nowhere/"},
		{"../root/../root/../root/", "/root/"},
		{"ужас/../кирлица/./също/може/да/е/път", "/кирлица/също/може/да/е/път/"},
	}

	for _, test := range tests {
		var got = parsePath(test.input_path)
		if got != test.expected_path {
			t.Errorf("Wanted `%s` but got `%s` for test case `%s`", test.expected_path,
				got, test.input_path)
			fmt.Printf("F")
		} else {
			fmt.Printf(".")
		}
	}

	fmt.Println("")
}
