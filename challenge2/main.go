package main

type turns map[bool]string

type DragonFractal struct {
	iteration int
	last      string
	rotations map[string]turns
}

func (dragon *DragonFractal) Next() string {
	var n = dragon.iteration

	if n == 0 {
		dragon.rotations = make(map[string]turns)

		dragon.rotations["up"] = make(turns)
		dragon.rotations["up"][true] = "left"
		dragon.rotations["up"][false] = "right"

		dragon.rotations["left"] = make(turns)
		dragon.rotations["left"][true] = "down"
		dragon.rotations["left"][false] = "up"

		dragon.rotations["down"] = make(turns)
		dragon.rotations["down"][true] = "right"
		dragon.rotations["down"][false] = "left"

		dragon.rotations["right"] = make(turns)
		dragon.rotations["right"][true] = "up"
		dragon.rotations["right"][false] = "down"

		dragon.iteration += 1
		dragon.last = "up"
		return dragon.last
	}

	var turn_left bool = !((((n & -(n)) << 1) & n) != 0)
	dragon.iteration += 1
	dragon.last = dragon.rotations[dragon.last][turn_left]
	return dragon.last
}
