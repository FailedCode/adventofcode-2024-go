package solvers

import (
	"fmt"
	"regexp"
	"strconv"
	// "time"
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
	robots := inputToRobots(input)

	// fmt.Printf("%v\n", robots[0])
	for i := 0; i < 100; i += 1 {
		moveRobots(robots, maxW, maxH)
	}
	// fmt.Printf("%v\n", robots[0])

	safetyFactor := calculateSafetyFactor(robots, maxW, maxH)

	return fmt.Sprintf("%v", safetyFactor)
}

func (s Day14Solver) Part2() string {

	maxW := 101
	maxH := 103

	input := utility.LoadInput(s.Day, s.InputSource)
	robots := inputToRobots(input)

	i := 0
	for true {
		i += 1
		moveRobots(robots, maxW, maxH)
		if robotsInARow(robots) {
			break;
		}
	}
	displayRobots(robots)

	return fmt.Sprintf("%v", i)
}

func inputToRobots(input []string) []Robot {
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
	return robots
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

func boundingBox(robots []Robot) int {
	minX := 99999
	maxX := 0
	minY := 99999
	maxY := 0
	for _, robot := range robots {
		if robot.px < minX {
			minX = robot.px
		}
		if robot.py < minY {
			minY = robot.py
		}
		if robot.px > maxX {
			maxX = robot.px
		}
		if robot.py > maxY {
			maxY = robot.py
		}
	}
	// fmt.Printf("boundingBox: (%v, %v), (%v, %v) \n", minX, minY, maxX, maxY)
	return (maxX - minX) * (maxY - minY)
}

func robotsInARow(robots []Robot) bool {
	field := [103][101]int{}
	for _, robot := range robots {
		field[robot.py][robot.px] += 1
	}
	adjecend := 0
	for _, row := range field {
		for _, value := range row {
			if value > 0 {
				adjecend += 1
			} else {
				adjecend = 0
			}
			if adjecend > 7 {
				return true
			}
		}
		adjecend = 0
	}
	return false
}

func displayRobots(robots []Robot) {
	display := [103][101]int{}
	for _, robot := range robots {
		display[robot.py][robot.px] += 1
	}
	for _, row := range display {
		for _, value := range row {
			if value > 0 {
				fmt.Printf("%v", value)
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}

}