package main

import (
	"fmt"
	"flag"
	"time"
	"log"
	"errors"
	"internal/solvers"
)

var day int
var part int

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
	flag.IntVar(&day, "day", 0, "Explicitly set the day to be solved")
	flag.IntVar(&part, "part", 0, "defaults to 0 (both)")
	flag.Parse()

	if day < 1 {
		// Use the current day of the month if none is specified.
		// Works great during december
		date := time.Now()
		day = date.Day()
	}
}

func getSolverByName(name string) (solvers.DaySolver, error) {
	// Compiled languages like go have no other options than
	// explicitly listing all Types that may be used, it seems
	var solvers = map[string]solvers.DaySolver {
		"Day1Solver": solvers.Day1Solver{},
		"Day2Solver": solvers.Day2Solver{},
	}
	solver := solvers[name]
	if solver == nil {
		return nil, errors.New("Unknown Solver")
	}
	return solvers[name], nil
}