package main

import "testing"
import "fmt"

func TestDragonFractalWithFirstFewLines(t *testing.T) {
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

func BenchmarkDragonFractal_Next(b *testing.B) {
	dragon := new(DragonFractal)
	for i := 0; i < b.N; i++ {
		dragon.Next()
	}
}

func BenchmarkDragonFractal_isNextTurnRight(b *testing.B) {
	dragon := new(DragonFractal)
	for i := 0; i < b.N; i++ {
		dragon.isNextTurnRight()
	}
}

func BenchmarkDragonFractal_translate(b *testing.B) {
	dragon := new(DragonFractal)
	for i := 0; i < b.N; i++ {
		dragon.translate(0)
		dragon.translate(1)
		dragon.translate(2)
		dragon.translate(3)
	}
}
