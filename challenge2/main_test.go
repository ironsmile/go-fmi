package main

import "testing"
import "fmt"

func Test(t *testing.T) {
	tests := []string{"up", "left", "down", "left", "down", "right", "down", "left"}

	dragon := new(DragonFractal)

	for index, expected := range tests {
		var got = dragon.Next()
		if got != expected {
			t.Errorf("Wanted `%s` but got `%s` on iteration `%d`", expected,
				got, index)
			fmt.Printf("F")
		} else {
			fmt.Printf(".")
		}
	}

	fmt.Println("")
}
