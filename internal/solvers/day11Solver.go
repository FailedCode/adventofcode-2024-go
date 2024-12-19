package solvers

import (
	"fmt"
	"strings"
	"strconv"
	"time"
	"internal/utility"
)

type Day11Solver struct {
	Day uint
	InputSource string
}

func (s Day11Solver) Part1() string {

	input := utility.LoadInput(s.Day, s.InputSource)
	stones := inputToIntSlice(input[0])

	// fmt.Printf("stones inital:\n%v\n", stones)
	for i := 0; i < 25; i += 1 {
		stones = changeStones(stones)
		// fmt.Printf("blink %v:\n%v\n", i, stones)
	}

	return fmt.Sprintf("%v", len(stones))
}

func (s Day11Solver) Part2() string {

	input := utility.LoadInput(s.Day, s.InputSource)
	stones := inputToIntSlice(input[0])

	//
	// Part 1 approach
	// blink 43: signal: killed :()
	//

	var stoneStack []Stone
	for _, n := range stones {
		stoneStack = append(stoneStack, Stone{v: n, generation: 75})
	}

	// fmt.Printf("stones inital:\n%v\n", stoneStack)
	stoneNumber := getStonesLineage(stoneStack, 1000)

	return fmt.Sprintf("%v", stoneNumber)
}

func inputToIntSlice(input string) []int {
	var stones []int
	numbers := strings.Split(input, " ")
	for _, number := range numbers {
		n, _ := strconv.Atoi(number)
		stones = append(stones, n)
	}
	return stones
}

func changeStones(stones []int) []int {
	var newStones []int
	for _, stone := range stones {
		if stone == 0 {
			// If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
			newStones = append(newStones, 1)
		} else if len(strconv.Itoa(stone)) % 2 == 0 {
    		// If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
    		left, right := splitNumber(stone)
			newStones = append(newStones, left, right)
    		// fmt.Printf("split stone \"%v\" => %v | %v\n", stone, left, right)
		} else {
    		// If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
			newStones = append(newStones, stone * 2024)
		}
	}
	return newStones
}

func splitNumber(i int) (int,int) {
	s := strconv.Itoa(i)
	l := len(s)
	half := l / 2
	left, _ := strconv.Atoi(s[0:half])
	right, _ := strconv.Atoi(s[half:l])
	return left, right
}

type Stone struct {
	v int
	generation int
}


// Let's track the linage of only one branch really deep
// and sum up the stones in that branch before throwing
// the branch away and starting the next branch
// ---
// This is easier on the memory but takes ages as well...
func getStonesLineage(stones []Stone, progress_ms int64) int64 {
	startTime := time.Now()
	timer := time.Now()
	var stone Stone
	var counter int64 = 0
	for len(stones) > 0 {
		stone, stones = pop(stones)
		for true {
			if progress_ms > 0 && time.Since(timer).Milliseconds() >= progress_ms {
				fmt.Printf("stack size: %v \tcounter: %v \tcalculating for: %v\n", len(stones), counter, time.Since(startTime))
				timer = time.Now()
			}

			if stone.generation == 0 {
				counter += 1
				break
			}

			if stone.v == 0 {

				// optimize the next few steps
				if stone.generation > 3 {
					// make value 1
					stone.generation -= 1
					// make 2024
					stone.generation -= 1
					// split value
					stones = append(stones, Stone{v: 24, generation: stone.generation -1})
	    			stone.v = 20
					stone.generation -= 1
					// this is all deterministic from here, so you could probably cache the next
					// X generations and skip all the computation
					continue
				}

				stone.v = 1
				stone.generation -= 1
			} else if len(strconv.Itoa(stone.v)) % 2 == 0 {
	    		left, right := splitNumber(stone.v)
				stones = append(stones, Stone{v: right, generation: stone.generation -1})
	    		stone.v = left
				stone.generation -= 1
			} else {
				stone.v *= 2024
				stone.generation -= 1
			}
		}
	}
	if progress_ms > 0 {
		fmt.Printf("duration: %v\n", time.Since(startTime))
	}
	return counter
}