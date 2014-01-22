package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAllStuck(t *testing.T) {
	set := make(map[[2][2]int]bool)
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			var stuck = [2][2]int{{x, y}, {-2, -2}}
			set[stuck] = true
		}
	}

	var allStuck = [4][4]rune{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
	}
	resultSet := make(map[[2][2]int]bool)
	for _, value := range playMall(allStuck) {
		resultSet[value] = true
	}

	if !reflect.DeepEqual(resultSet, set) {
		t.Fail()
	}
}

func TestSingleDweller(t *testing.T) {
	var singleDweller = [4][4]rune{
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', 'X'},
	}
	result := playMall(singleDweller)

	j := 0
	for i := 103; i > 3; i-- {
		start := i % 4
		end := (i - 1) % 4
		var currentElement = [2][2]int{{start, start}, {end, end}}
		if !reflect.DeepEqual(result[j], currentElement) {
			t.Fail()
		}
		j++
	}

	if j != len(result)-1 {
		t.Fail()
	}

	var currentElement = [2][2]int{{3, 3}, {-1, -1}}
	if !reflect.DeepEqual(result[j], currentElement) {
		t.Fail()
	}
}

func TestEmptyMall(t *testing.T) {
	var singleDweller = [4][4]rune{
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
		{'-', '-', '-', '-'},
	}
	result := playMall(singleDweller)

	if len(result) != 0 {
		t.Errorf("There were ghosts in the empty mall!")
	}
}

func TestPossibleMoves(t *testing.T) {
	check := func(tested [2]int, calculated [][2]int, expected [][2]int) {
		if len(calculated) != len(expected) {
			t.Errorf("Pos %#v. Expected %d results but there were %d", tested,
				len(expected), len(calculated))
			return
		}

		for ind, pos := range expected {
			if pos[0] != calculated[ind][0] || pos[1] != calculated[ind][1] {
				t.Errorf("Pos %#v. Expected %#v but got %#v for index %d", tested,
					pos, calculated[ind], ind)
			}
		}
	}

	pos := [2]int{1, 1}
	calculated := possibleMoves(pos)
	expected := [][2]int{{0, 0}, {2, 2}, {0, 2}, {2, 0}}
	check(pos, calculated, expected)

	pos = [2]int{0, 0}
	calculated = possibleMoves(pos)
	expected = [][2]int{{3, 3}, {1, 1}, {3, 1}, {1, 3}}
	check(pos, calculated, expected)

	pos = [2]int{3, 3}
	calculated = possibleMoves(pos)
	expected = [][2]int{{2, 2}, {0, 0}, {2, 0}, {0, 2}}
	check(pos, calculated, expected)

	pos = [2]int{1, 1}
	calculated = possibleMoves(pos)
	expected = [][2]int{{0, 0}, {2, 2}, {0, 2}, {2, 0}}
	check(pos, calculated, expected)
}

func TestNonInteractingDwellers(t *testing.T) {
	var NonInteractingDwellers = [4][4]rune{
		{'-', '-', '-', 'X'},
		{'-', '-', '-', 'X'},
		{'-', '-', '-', 'X'},
		{'-', '-', '-', 'X'},
	}
	result := playMall(NonInteractingDwellers)

	if 404 != len(result) {
		t.Fail()
	}

	for i := 0; i < len(result); i++ {
		// In the last move, dwellers must be at the position where they started from
		if reflect.DeepEqual(result[i][1], [2]int{-1, -1}) {
			if result[i][0][1] != 3 {
				t.Fail()
			}
			continue
		}

		// Check if all of them move to up-right
		if !reflect.DeepEqual(result[i][0], [2]int{(result[i][1][0] + 1) % 4,
			(result[i][1][1] + 1) % 4}) {
			t.Fail()
		}
	}
}

func TestDifferentReplays(t *testing.T) {
	var malls = [][4][4]rune{
		{
			{'-', '-', '-', '-'},
			{'-', '-', '-', '-'},
			{'-', '-', 'X', '-'},
			{'-', '-', '-', '-'},
		},
		{
			{'X', '-', 'X', '-'},
			{'-', '-', '-', '-'},
			{'X', '-', 'X', '-'},
			{'-', '-', '-', '-'},
		},
		{
			{'X', '-', 'X', '-'},
			{'-', 'X', '-', '-'},
			{'X', '-', 'X', '-'},
			{'-', '-', '-', '-'},
		},
		{
			{'X', 'X', 'X', '-'},
			{'X', '-', 'X', '-'},
			{'X', 'X', 'X', '-'},
			{'-', '-', '-', '-'},
		},
		{
			{'-', '-', '-', '-'},
			{'X', 'X', 'X', 'X'},
			{'-', 'X', 'X', '-'},
			{'-', '-', '-', '-'},
		},
		{
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
		},
		{
			{'X', '-', 'X', '-'},
			{'-', 'X', '-', 'X'},
			{'X', '-', 'X', '-'},
			{'-', 'X', '-', 'X'},
		},
		{
			{'X', 'X', 'X', 'X'},
			{'-', '-', '-', '-'},
			{'X', 'X', 'X', 'X'},
			{'-', '-', '-', '-'},
		},
		{
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'-', '-', '-', '-'},
			{'-', '-', '-', '-'},
		},
		{
			{'X', 'X', 'X', '-'},
			{'X', '-', '-', '-'},
			{'X', 'X', 'X', '-'},
			{'X', 'X', 'X', '-'},
		},
		{
			{'X', 'X', 'X', '-'},
			{'X', '-', 'X', '-'},
			{'X', '-', 'X', '-'},
			{'X', 'X', 'X', '-'},
		},
		{
			{'X', '-', '-', '-'},
			{'X', '-', '-', '-'},
			{'X', '-', '-', '-'},
			{'X', 'X', 'X', '-'},
		},
		{
			{'-', 'X', '-', '-'},
			{'X', '-', 'X', '-'},
			{'X', 'X', 'X', '-'},
			{'X', '-', 'X', '-'},
		},
		{
			{'X', '-', 'X', '-'},
			{'X', 'X', 'X', '-'},
			{'X', '-', 'X', '-'},
			{'X', '-', 'X', '-'},
		},
		{
			{'X', 'X', 'X', '-'},
			{'X', '-', '-', '-'},
			{'X', 'X', 'X', '-'},
			{'X', 'X', 'X', '-'},
		},
		{
			{'X', '-', '-', '-'},
			{'X', '-', '-', '-'},
			{'-', '-', '-', '-'},
			{'X', '-', '-', '-'},
		},
		{
			{'-', '-', '-', '-'},
			{'-', '-', '-', '-'},
			{'-', '-', '-', '-'},
			{'-', '-', '-', '-'},
		},
	}

	for index, mall := range malls {
		t.Logf("Testing with mall %d", index)
		result := playMall(mall)
		validateReplay(t, mall, result)
	}
}

type testDweller struct {
	moves   int
	stopped bool
	stuck   bool
}

func validateReplay(t *testing.T, mall [4][4]rune, replay [][2][2]int) {

	var board [4][4]*testDweller

	numberOfDwellers := 0
	for mallX, coll := range mall {
		for mallY, cell := range coll {
			if cell == '-' {
				continue
			}
			board[mallX][mallY] = new(testDweller)
			numberOfDwellers += 1
		}
	}

	occupied := func(pos [2]int) bool {
		return board[pos[0]][pos[1]] != nil
	}

	fatal := func(msg string) {
		debugMessage := fmt.Sprintf(`%s
		Mall starting possition:
		%#v
		All Moves: %#v
		`, msg, mall, replay)
		t.Fatal(debugMessage)
	}

	for moveIndex, move := range replay {
		from := move[0]
		to := move[1]

		if !occupied(from) {
			fatal(fmt.Sprintf("Trying to move nondweller on move %d",
				moveIndex))
		}

		dweller := board[from[0]][from[1]]
		possibleShops := possibleMoves(from)

		if dweller.stopped {
			fatal(fmt.Sprintf("Moving stopped dweller on move %d", moveIndex))
		}

		if to == [2]int{-2, -2} {
			for _, shop := range possibleShops {
				if !occupied(shop) {
					fatal(fmt.Sprintf("Says stuck when it is not on move %d",
						moveIndex))
				}
			}
			dweller.stopped = true
			dweller.stuck = true
			continue
		}

		if to == [2]int{-1, -1} {
			if dweller.moves != 100 {
				fatal(fmt.Sprintf("Says end of turns when there are left on move %d",
					moveIndex))
			}
			dweller.stopped = true
			continue
		}

		if occupied(to) {
			fatal(fmt.Sprintf("Trying to move in occupied shop on move %d",
				moveIndex))
		}

		dweller.moves += 1

		board[from[0]][from[1]] = nil
		board[to[0]][to[1]] = dweller
	}

	foundDwellers := 0
	for mallX, coll := range board {
		for mallY, cell := range coll {
			if cell == nil {
				continue
			}
			dweller := board[mallX][mallY]
			foundDwellers += 1

			if !dweller.stopped {
				fatal("A dweller has not finished its turns")
			}

			if dweller.stuck {
				continue
			}

			if dweller.moves != 100 {
				fatal("A dweller stopped before its 100th turn")
			}
		}
	}

	if foundDwellers != numberOfDwellers {
		fatal("verifyReplay lost some dwellers")
	}
}
