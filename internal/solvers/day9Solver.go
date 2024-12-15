package solvers

import (
	"fmt"
	"strings"
	"strconv"
	"internal/utility"
)

type Day9Solver struct {
	Day uint
	InputSource string
}

func (s Day9Solver) Part1() string {

	input := utility.LoadInput(s.Day, s.InputSource)
	decompressed := decompress(input[0])
	defragmented := defragment(decompressed)
	checksum := calculateChecksum(defragmented)

	return fmt.Sprintf("%v", checksum)
}

func (s Day9Solver) Part2() string {
	return fmt.Sprintf("todo: implement Part2")
}


func decompress(input string) []int {
	var decompressed []int
	var value int = 0
	chars := strings.Split(input, "")
	for key, char := range chars {
		length, _ := strconv.Atoi(char)
		if key % 2 == 0 {
			file := make([]int, length)
			fill(file, value)
			decompressed = append(decompressed, file...)
			value += 1
			// fmt.Printf("file: %v => %v\n", value, length)
		} else {
			empty := make([]int, length)
			fill(empty, -1)
			decompressed = append(decompressed, empty...)
			// fmt.Printf("empty: %v => %v\n", ".", length)
		}
	}
	return decompressed
}

func fill[T any](slice []T, val T) {
	for i := range slice {
		slice[i] = val
	}
}


func defragment(fs []int) []int {
	for position, value := range fs {
		if value != -1 {
			continue
		}

		var moveValue int
		moveValue, fs = pop(fs)
		for moveValue == -1 {
			moveValue, fs = pop(fs)
		}

		// slice will be shortend, so check if we are done
		if position >= len(fs) {
			fs = append(fs, moveValue)
			break
		}

		fs[position] = moveValue
		// fmt.Printf("now override %v with %v\nfs looks like:\n%v\n", fs[position], moveValue, fs)
	}
	return fs
}

// remove the last element of the slice and return it
func pop[T any](slice []T) (T, []T) {
	p := len(slice)-1
	v := slice[p]
	slice = slice[:p]
	return v, slice
}

func calculateChecksum(fs []int) int {
	var sum int = 0
	for position, value := range fs {
		sum += position * value
	}
	return sum
}