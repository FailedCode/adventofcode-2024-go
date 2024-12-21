package solvers

import (
	"fmt"
	"regexp"
	"strconv"
	"internal/utility"
)

type Day14Solver struct {
	Day uint
	InputSource string
}

type Robot struct {
	px int
	py int
	vx int
	vy int
}

func (s Day14Solver) Part1() string {

	// real
	maxW := 101
	maxH := 103
	if s.InputSource != "day" {
		// example
		maxW = 11
		maxH = 7
	}

	input := utility.LoadInput(s.Day, s.InputSource)
	r := regexp.MustCompile(`p=(\d+),(\d+)\sv=(-?\d+),(-?\d+)`)
	var robots []Robot
	for _, line := range input {
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			px, _ := strconv.Atoi(match[1])
			py, _ := strconv.Atoi(match[2])
			vx, _ := strconv.Atoi(match[3])
			vy, _ := strconv.Atoi(match[4])
			robots = append(robots, Robot{px: px, py: py, vx: vx, vy: vy})
		}
	}

	// fmt.Printf("%v\n", robots[0])
	for i := 0; i < 100; i += 1 {
		moveRobots(robots, maxW, maxH)
	}
	// fmt.Printf("%v\n", robots[0])

	safetyFactor := calculateSafetyFactor(robots, maxW, maxH)

	return fmt.Sprintf("%v", safetyFactor)
}

func (s Day14Solver) Part2() string {
	return fmt.Sprintf("todo: implement Part2")
}

func moveRobots(robots []Robot, maxW int, maxH int) {
	for i := range robots {
		robots[i].px = (robots[i].px + robots[i].vx + maxW) % maxW
		robots[i].py = (robots[i].py + robots[i].vy + maxH) % maxH
	}
}

func calculateSafetyFactor(robots []Robot, maxW int, maxH int) int {
	maxW2 := maxW/2
	maxH2 := maxH/2

	q1minX := 0
	q1minY := 0
	q1maxX := maxW2-1
	q1maxY := maxH2-1

	q2minX := maxW2+1
	q2minY := 0
	q2maxX := maxW-1
	q2maxY := maxH2-1

	q3minX := 0
	q3minY := maxH2+1
	q3maxX := maxW2-1
	q3maxY := maxH-1

	q4minX := maxW2+1
	q4minY := maxH2+1
	q4maxX := maxW-1
	q4maxY := maxH-1

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0
	for _, robot := range robots {
		if robot.IsInRect(q1minX,q1minY,q1maxX,q1maxY) {
			q1 += 1
		} else if robot.IsInRect(q2minX,q2minY,q2maxX,q2maxY) {
			q2 += 1
		} else if robot.IsInRect(q3minX,q3minY,q3maxX,q3maxY) {
			q3 += 1
		} else if robot.IsInRect(q4minX,q4minY,q4maxX,q4maxY) {
			q4 += 1
		}
	}
	// fmt.Printf("quadrants: %v, %v, %v, %v \n", q1, q2, q3, q4)
	return q1 * q2 * q3 * q4
}

func (r Robot) IsInRect(minX int, minY int, maxX int, maxY int) bool {
	if r.px < minX {return false}
	if r.py < minY {return false}
	if r.px > maxX {return false}
	if r.py > maxY {return false}
	return true
}