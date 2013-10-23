package main

const (
	up    = "up"
	down  = "down"
	right = "right"
	left  = "left"
)

type turns map[bool]string

var rotations = map[string]turns{
	"up": turns{
		true:  "left",
		false: "right",
	},

	"left": turns{
		true:  "down",
		false: "up",
	},

	"down": turns{
		true:  "right",
		false: "left",
	},

	"right": turns{
		true:  "up",
		false: "down",
	},
}

type DragonFractal struct {
	iteration int
	last      string
}

func (dragon *DragonFractal) Next() string {
	var n = dragon.iteration

	if n == 0 {
		dragon.iteration += 1
		dragon.last = up
		return dragon.last
	}

	var turn_left bool = !((((n & -(n)) << 1) & n) != 0)
	dragon.iteration += 1
	dragon.last = rotations[dragon.last][turn_left]
	return dragon.last
}
