package solvers

import (
	"fmt"
	"strings"
	"strconv"
	"internal/utility"
)

type Day2Solver struct {}

func (s Day2Solver) Part1() string {

	input := utility.LoadInput(2)

	safeReports := 0
	for _, line := range input {
		if len(line) == 0 {
			break
		}
		report := lineToNumberSlice(line)
		// fmt.Println(report)
		if isReportSafe(report) {
			safeReports += 1
		}
	}

	return fmt.Sprintf("%v", safeReports)
}

func (s Day2Solver) Part2() string {
	return "todo: implement Part2"
}

func isReportSafe(report []int) bool {
	last_n := report[0]
	sign := 0
	for _, n := range report[1:] {
		diff := utility.Abs(last_n - n)
		if diff < 1 || diff > 3 {
			// fmt.Printf("diff: %v - %v = %v\n", last_n, n, diff)
			return false
		}
		newSign := utility.Sgn(last_n - n)
		if sign != 0 && newSign != sign {
			// fmt.Printf("direction change: %v => %v (%v => %v)\n", sign, newSign, last_n, n)
			return false
		}
		sign = newSign
		last_n = n
	}

	return true
}

func lineToNumberSlice(line string) []int {
	s := strings.Split(line, " ")
	integers := make([]int, 0, len(line))
	for _, x := range s {
		n, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		integers = append(integers, n)
	}
	return integers
}