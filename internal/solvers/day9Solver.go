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

	input := utility.LoadInput(s.Day, s.InputSource)
	decompressed := decompress(input[0])
	defragmented := defragment_blockwise(decompressed)
	checksum := calculateChecksum(defragmented)

	return fmt.Sprintf("%v", checksum)
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
		if value < 0 {
			continue
		}
		sum += position * value
	}
	return sum
}

func defragment_blockwise(fs []int) []int {
	file_checked := make(map[int]bool)
	maxPosition := len(fs)-1
	filePosition := maxPosition
	fileStartPos := maxPosition
	fileLength := 0
	fileValue := 0

	emptyPosition := 0
	emptyStartPos := 0
	emptyLength := 0
	emptyValue := 0

	// fmt.Printf("%v\n", fs)

	for true {
		// find the next file block from the right
		fileValue = fs[filePosition]
		for fileValue == -1 {
			filePosition -= 1
			fileValue = fs[filePosition]
		}

		// We have reached the beginning of the filesystem
		if fileValue == 0 {
			break
		}

		fileStartPos = filePosition
		// now fileValue is >-1 : count how long the block is
		for fs[filePosition - fileLength] == fileValue {
			fileLength += 1
		}

		// file was checked before? move past it
		if file_checked[fileValue] {
			filePosition -= fileLength
			fileLength = 0
			continue
		}

		// we will check it now - if it moves or not, doesn't matter
		file_checked[fileValue] = true

		for true {
			// for this file, try to find a free space from the left:
			emptyValue = fs[emptyPosition]
			for emptyValue != -1 {
				emptyPosition += 1
				emptyValue = fs[emptyPosition]
				if emptyPosition >= fileStartPos {
					break
				}
			}
			emptyStartPos = emptyPosition

			// start of empty block found: how long is it?
			for emptyValue == -1 {
				emptyPosition += 1
				emptyValue = fs[emptyPosition]
				if emptyPosition >= fileStartPos {
					break
				}
			}
			emptyLength = emptyPosition - emptyStartPos

			// is the empty space after the current file: try the next file
			if emptyStartPos + emptyLength -1 > fileStartPos - fileLength {
				emptyPosition = 0
				filePosition -= fileLength
				fileLength = 0
				break
			}

			// not enough space to put the file here
			if emptyLength < fileLength {
				// fmt.Printf("Does not fit, resume search!\n")
				continue
			}

			// fmt.Printf("========================================\n")
			// fmt.Printf("Fits in here!\n")
			// fmt.Printf("%v\n", fs)
			// fmt.Printf("block %v size: %v starting at: %v) => start %v len: %v \n", fileValue, fileLength, fileStartPos, emptyStartPos, emptyLength)

			// there is enough space: Move the file
			i := 0
			for fileLength > i {
				fs[emptyStartPos + i] = fileValue
				fs[fileStartPos - i] = -1
				i += 1
			}
			// start with the next file
			emptyPosition = 0
			filePosition -= fileLength
			fileLength = 0
			// fmt.Printf("moved file %v\n", fileValue)
			break
		}
	}
	return fs
}