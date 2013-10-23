// Here we create a DragonFractal type which can be used to guide you
// how to draw a dragon curve. Its only method tells you in what direction
// you should draw the next line

package main

// right MUST be the first in this enum so that its value is 0.
// We are depending on this to remove the if dragon.iteration == 0 on every call
// from the Next method.
const (
	right = iota
	down
	up
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

// Use this type to draw a dragon curve
type DragonFractal struct {
	iteration int // how many lines have we drawn already
	last      int // what direction was the last line
}

// Tells you in which direction to draw next
func (dragon *DragonFractal) Next() string {
	dragon.last = rotations[dragon.last][dragon.isNextTurnRight()]
	dragon.iteration++
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
