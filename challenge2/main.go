package main

type turns map[bool]string

type DragonFractal struct {
	iteration int
	last      string
}

func (dragon *DragonFractal) Next() string {
	var n = dragon.iteration

	if n == 0 {
		dragon.iteration += 1
		dragon.last = "up"
		return dragon.last
	}

	var turn_left bool = !((((n & -(n)) << 1) & n) != 0)

	rotations := make(map[string]turns)

	rotations["up"] = make(turns)
	rotations["up"][true] = "left"
	rotations["up"][false] = "right"

	rotations["left"] = make(turns)
	rotations["left"][true] = "down"
	rotations["left"][false] = "up"

	rotations["down"] = make(turns)
	rotations["down"][true] = "right"
	rotations["down"][false] = "left"

	rotations["right"] = make(turns)
	rotations["right"][true] = "up"
	rotations["right"][false] = "down"

	dragon.iteration += 1
	dragon.last = rotations[dragon.last][turn_left]
	return dragon.last
}
