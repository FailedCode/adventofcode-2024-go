package solvers

import (
	"fmt"
	"regexp"
	"strconv"
	"internal/utility"
)

type Day3Solver struct {
	Day uint
	InputSource string
}

func (s Day3Solver) Part1() string {
	input := utility.LoadInput(s.Day, s.InputSource)
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	var result int = 0
	for _, line := range input {
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			n1, _ := strconv.Atoi(match[1])
			n2, _ := strconv.Atoi(match[2])
			result += n1 * n2
		}
	}

	return fmt.Sprintf("%v", result)
}

func (s Day3Solver) Part2() string {
	input := utility.LoadInput(s.Day, s.InputSource)
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	var result int = 0
	var sum_enabled bool = true
	for _, line := range input {
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[1] != "" && match[2] != "" {
				if sum_enabled {
					n1, _ := strconv.Atoi(match[1])
					n2, _ := strconv.Atoi(match[2])
					result += n1 * n2
				}
			} else {
				if match[0] == "do()" {
					sum_enabled = true
				}
				if match[0] == "don't()" {
					sum_enabled = false
				}
			}
		}
	}

	return fmt.Sprintf("%v", result)
}