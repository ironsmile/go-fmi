package main

const (
	up = iota
	down
	right
	left
)

type turns map[bool]int

var rotations = map[int]turns{
	up: turns{
		true:  right,
		false: left,
	},

	left: turns{
		true:  up,
		false: down,
	},

	down: turns{
		true:  left,
		false: right,
	},

	right: turns{
		true:  down,
		false: up,
	},
}

type DragonFractal struct {
	iteration int
	last      int
}

func (dragon *DragonFractal) Next() string {

	if dragon.iteration == 0 {
		dragon.iteration += 1
		dragon.last = up
		return dragon.translate(up)
	}

	var turn_left bool = dragon.isNextTurnRight()
	dragon.iteration += 1
	dragon.last = rotations[dragon.last][turn_left]
	return dragon.translate(dragon.last)
}

func (dragon *DragonFractal) translate(direction int) string {
	switch direction {
	case up:
		return "up"
	case down:
		return "down"
	case left:
		return "left"
	case right:
		return "right"
	}
	return ""
}

func (dragon *DragonFractal) isNextTurnRight() bool {
	return (((dragon.iteration & -(dragon.iteration)) << 1) & dragon.iteration) != 0
}
