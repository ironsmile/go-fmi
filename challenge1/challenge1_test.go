package challenge1

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		program string
		wanted  bool
	}{
		{``, true},
		{`я`, false},
		{`
    package main

    import "fmt"

    func main() {
        баба := 1
    }
`, false},
		{`package main

import "fmt"

func is_chinese() bool {
    return "汉语" == "漢語"
}

func main () {
    fmt.Printf("Well, it is %b", is_chinese())
}`, false},
		{`// An empty struct.
struct {}

// A struct with 6 fields.
struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}`, true},
		{`const a = 2 + 3.0          // a == 5.0   (untyped floating-point constant)
const b = 15 / 4           // b == 3     (untyped integer constant)
const c = 15 / 4.0         // c == 3.75  (untyped floating-point constant)
const Θ float64 = 3/2      // Θ == 1.0   (type float64, 3/2 is integer division)
const Π float64 = 3/2.     // Π == 1.5   (type float64, 3/2. is float division)
const d = 1 << 3.0         // d == 8     (untyped integer constant)
const e = 1.0 << 3         // e == 8     (untyped integer constant)
const f = int32(1) << 33   // illegal    (constant 8589934592 overflows int32)
const g = float64(2) >> 1  // illegal    (float64(2) is a typed floating-point constant)
const h = "foo" > "bar"    // h == true  (untyped boolean constant)
const j = true             // j == true  (untyped boolean constant)
const k = 'w' + 1          // k == 'x'   (untyped rune constant)
const l = "hi"             // l == "hi"  (untyped string constant)
const m = string(k)        // m == "x"   (type string)
const Σ = 1 - 0.707i       //            (untyped complex constant)
const Δ = Σ + 2.0e-4       //            (untyped complex constant)
const Φ = iota*1i - 1/1i   //            (untyped complex constant)`, false},
		{`Strictly speaking this is not a program but hey! It helps nevertheless.`, true},
		{`~!@#$%^&*()_-=+|\/`, true},
	}
	for _, test_tuple := range tests {
		got := hasOnlyLatinSymbols(test_tuple.program)
		if got != test_tuple.wanted {
			t.Errorf("Wanted %b but got %b for \n%s", test_tuple.wanted, got, test_tuple.program)
		}
	}
}
