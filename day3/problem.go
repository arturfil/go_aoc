package day3

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/arturfil/go_aoc/helpers"
)

func PartOne() {
	file, _ := helpers.ReadInput("./day3/input.txt")
    
	pattern := `(do\(\)|don't\(\))|mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)

    var matches [][]string 
    for _, line := range file {
        match := re.FindAllStringSubmatch(line, -1)
        fmt.Println("match -> \n", match)
        matches = append(matches, match...)
    }

    for _, match := range matches {
       fmt.Println("match", match) 
    }

    totScore := calculateScore(matches, false)
    fmt.Println("Tot Score:", totScore)
}

func PartTwo() {
	file, _ := helpers.ReadInput("./day3/input.txt")

	validPattern := `(do\(\)|don't\(\))|mul\((\d+),(\d+)\)`
    re := regexp.MustCompile(validPattern)

    // find all matches in sample
    var matches [][]string
    for _, line := range file {
        match := re.FindAllStringSubmatch(line, -1)
        matches = append(matches, match...)
    }

    totScore := calculateScore(matches, true)
    fmt.Println("Tot Score:", totScore)
}

func calculateScore(matches [][]string, isControlled bool) int {
	var totScore int
    enabled := true

	for _, match := range matches {

        if match[1] == "don't()" && isControlled {
            enabled = false
            continue
        }
        if match[1] == "do()" && isControlled {
            enabled = true
            continue
        }

        if enabled {
            n1, _ := strconv.Atoi(match[2])
            n2, _ := strconv.Atoi(match[3])

            totScore += n1 * n2
        }
	}

	return totScore
}
