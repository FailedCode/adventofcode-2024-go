package solvers

import (
	"fmt"
	"internal/utility"
	"regexp"
	"strconv"
	"slices"
)

type Day1Solver struct {}

func (s Day1Solver) Part1() string {

	input := utility.LoadInput(1)
    leftNumbers := []int{}
    rightNumbers := []int{}
    // raw string, so slashes do not need to be escaped
    r := regexp.MustCompile(`(\d+)\s+(\d+)`)

	for _, line := range input {
		if len(line) == 0 {
			break
		}
		elements := r.FindStringSubmatch(line)
		e1, _ := strconv.Atoi(elements[1])
		leftNumbers = append(leftNumbers, e1)
		e2, _ := strconv.Atoi(elements[2])
		rightNumbers = append(rightNumbers, e2)
	}
	slices.Sort(leftNumbers)
	slices.Sort(rightNumbers)

	var distance int = 0
	for i, lValue := range leftNumbers {
		rValue := rightNumbers[i]
		distance += utility.Abs(lValue - rValue)
	}

	return fmt.Sprintf("%v", distance)
}

func (s Day1Solver) Part2() string {
	return "todo: implement Part2"
}
