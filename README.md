# Advent of Code 2024
https://adventofcode.com/2024

Solutions in [GO](https://go.dev/learn/)

## How to run

* Run the current day: `go run .`
* Run a specific day: `go run . --day=<day>` e.g. `go run . --day=6`
* Run a specific part: `go run . --part=<1,2>` e.g. `go run . --day=2 --part=2`
* Run a different input: `go run . --source=<name>` e.g. `go run . --day=2 --part=2 --source=example`

## GO learnings

* There are no default values for function parameters
* string to int conversion requieres explicit `strconv.Atoi`
* `math.Abs` is only implemented for floats