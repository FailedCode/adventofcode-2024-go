package main

import (
	"fmt"
	"flag"
	"time"
	"log"
	"errors"
	"internal/solvers"
)

var day uint
var part uint
var inputSource string

func main() {
	parseArguments()

	solver, err := getSolverByName(fmt.Sprintf("Day%dSolver", day))
	if err != nil {
		log.Fatalln(err)
	}

	if part == 0 || part == 1 {
		fmt.Println(fmt.Sprintf("Day %v, Part 1:", day))
		fmt.Println(solver.Part1())
	}
	if part == 0 || part == 2 {
		fmt.Println(fmt.Sprintf("Day %v, Part 2:", day))
		fmt.Println(solver.Part2())
	}
}

func parseArguments() {
	flag.UintVar(&day, "day", 0, "Explicitly set the day to be solved")
	flag.UintVar(&part, "part", 0, "defaults to 0 (both)")
	flag.StringVar(&inputSource, "source", "day", "selects what input file prefix to load, defaults to \"day\"")
	flag.Parse()

	if day < 1 {
		// Use the current day of the month if none is specified.
		// Works great during december
		day = uint(time.Now().Day())
	}
}

func getSolverByName(name string) (solvers.DaySolver, error) {
	// Compiled languages like go have no other options than
	// explicitly listing all Types that may be used, it seems
	var solver solvers.DaySolver = nil
	if name == "Day1Solver" {
		solver = &solvers.Day1Solver{Day: day, InputSource: inputSource}
	}
	if name == "Day2Solver" {
		solver = &solvers.Day2Solver{Day: day, InputSource: inputSource}
	}
	if name == "Day3Solver" {
		solver = &solvers.Day3Solver{Day: day, InputSource: inputSource}
	}
	if solver == nil {
		return nil, errors.New("Unknown Solver")
	}
	return solver, nil
}