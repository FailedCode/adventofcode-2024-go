package solvers

import (
	"fmt"
	"strings"
	"strconv"
	"internal/utility"
)

type Day2Solver struct {
	Day uint
	InputSource string
}

func (s Day2Solver) Part1() string {

	input := utility.LoadInput(s.Day, s.InputSource)

	safeReports := 0
	for _, line := range input {
		if len(line) == 0 {
			break
		}
		report := lineToNumberSlice(line)
		// fmt.Println(report)
		if isReportSafe(report, 0) {
			safeReports += 1
		}
	}

	return fmt.Sprintf("%v", safeReports)
}

func (s Day2Solver) Part2() string {

	input := utility.LoadInput(s.Day, s.InputSource)

	safeReports := 0
	for _, line := range input {
		if len(line) == 0 {
			break
		}
		report := lineToNumberSlice(line)
		// fmt.Println(report)
		if isReportSafe(report, 1) {
			safeReports += 1
		} else if isReportSafe(report[1:], 0) {
			// in case the report is something like [1 3 2 1]
			safeReports += 1
		}
	}

	return fmt.Sprintf("%v", safeReports)
}

func isReportSafe(report []int, dampening int) bool {
	last_n := report[0]
	sign := 0
	for position, n := range report[1:] {
		diff := utility.Abs(last_n - n)
		if diff < 1 || diff > 3 {
			// fmt.Printf("diff: %v - %v = %v\n", last_n, n, diff)
			if dampening > 0 {
				fixedReport := make([]int, 0, len(report)-1)
				fixedReport = append(fixedReport, report[:position+1]...)
				fixedReport = append(fixedReport, report[position+2:]...)
				return isReportSafe(fixedReport, dampening-1)
			}
			return false
		}
		newSign := utility.Sgn(last_n - n)
		if sign != 0 && newSign != sign {
			// fmt.Printf("direction change: %v => %v (%v => %v)\n", sign, newSign, last_n, n)
			if dampening > 0 {
				fixedReport := make([]int, 0, len(report)-1)
				fixedReport = append(fixedReport, report[:position+1]...)
				fixedReport = append(fixedReport, report[position+2:]...)
				return isReportSafe(fixedReport, dampening-1)
			}
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