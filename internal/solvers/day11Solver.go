package solvers

import (
	"fmt"
	"strings"
	"strconv"
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
	return fmt.Sprintf("todo: implement Part2")
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