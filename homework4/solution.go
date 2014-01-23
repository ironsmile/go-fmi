package main

import "fmt"
import "sync"

var (
	OCCUPIED = 'X'
	EMPTY    = '-'
)

var (
	END_OF_TURNS = [2]int{-1, -1}
	STUCK        = [2]int{-2, -2}
)

func possibleMoves(position [2]int) [][2]int {

	var out [][2]int

	for _, diff := range [][2]int{{-1, -1}, {1, 1}, {1, -1}, {-1, 1}} {
		x := (position[0] + diff[0] + 4) % 4
		y := (position[1] + diff[1] + 4) % 4
		out = append(out, [2]int{x, y})
	}

	return out
}

type Mall struct {
	// The current state of the mall
	world [4][4]rune

	// Lock representing each shop
	locks [4][4]sync.Mutex

	// To make sure there is only one goroutine which is trying to lock something
	// in the world.
	gettinLocks sync.Mutex
}

/*
This Lock method locks sequentially. This makes it possible for one goroutine
to acquire lock A and start waiting for lock B while an other routine acquire B and
wait for A - deadlock. To fight this problem we have the gettingLocks lock which
makes sure only one of our goroutines is locking at the moment.
*/
func (mall *Mall) Lock(positions [][2]int) {
	mall.gettinLocks.Lock()
	for _, pos := range positions {
		mall.locks[pos[0]][pos[1]].Lock()
	}
	mall.gettinLocks.Unlock()
}

func (mall *Mall) Unlock(positions [][2]int) {
	for _, pos := range positions {
		mall.locks[pos[0]][pos[1]].Unlock()
	}
}

// To be only used with locked from and to
func (mall *Mall) Move(from [2]int, to [2]int) error {
	if mall.Occupied(to) || !mall.Occupied(from) {
		return fmt.Errorf("Invalid move from %#v to %#v", from, to)
	}
	mall.world[from[0]][from[1]] = EMPTY
	mall.world[to[0]][to[1]] = OCCUPIED
	return nil
}

// To be only used with locked pos
func (mall *Mall) Occupied(pos [2]int) bool {
	return mall.world[pos[0]][pos[1]] == OCCUPIED
}

// Our main function
func playMall(mallStart [4][4]rune) [][2][2]int {
	var allMoves [][2][2]int

	wg := new(sync.WaitGroup)

	movesChan := make(chan [2][2]int)

	mall := new(Mall)
	mall.world = mallStart

	maller := func(position [2]int) {
		defer wg.Done()

		for i := 0; i < 100; i++ {
			moved := false
			allPossibleMoves := possibleMoves(position)

			var locking [][2]int
			locking = append(allPossibleMoves, position)
			mall.Lock(locking)

			for _, nextPosition := range allPossibleMoves {
				err := mall.Move(position, nextPosition)
				if err == nil {
					moved = true
					movesChan <- [2][2]int{position, nextPosition}
					position = nextPosition
					break
				}
			}

			if moved {
				mall.Unlock(locking)
				continue
			}

			movesChan <- [2][2]int{position, STUCK}
			mall.Unlock(locking)
			return
		}

		movesChan <- [2][2]int{position, END_OF_TURNS}
	}

	for mallX, coll := range mall.world {
		for mallY, cell := range coll {
			if cell == EMPTY {
				continue
			}
			wg.Add(1)
			go maller([2]int{mallX, mallY})
		}
	}

	aggw := new(sync.WaitGroup)
	aggw.Add(1)
	go func() {
		defer aggw.Done()
		for move := range movesChan {
			allMoves = append(allMoves, move)
		}
	}()

	wg.Wait()
	close(movesChan)
	aggw.Wait()

	return allMoves
}
