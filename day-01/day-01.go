package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const CIRCLE_SIZE = 100

func main() {
	file, err := os.Open("day-01/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	rotations := strings.Split(strings.TrimSpace(string(data)), "\n")
	fmt.Println(solve(rotations))
}

func solve(rotations []string) int {
	position, zeroCount := 50, 0

	for _, rotation := range rotations {
		direction := rotation[0]
		degrees, err := strconv.Atoi(rotation[1:])
		if err != nil {
			panic(err)
		}
		if direction == 'L' {
			degrees *= -1
		}
		position = ((position+degrees)%CIRCLE_SIZE + CIRCLE_SIZE) % CIRCLE_SIZE

		if position == 0 {
			zeroCount++
		}
	}
	return zeroCount
}
