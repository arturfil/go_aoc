package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/arturfil/go_aoc/helpers"
)

/*
var tests [][]int = [][]int {
    {7, 6, 4, 2, 1},
    // {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1},
    // {1, 3, 2, 4, 5},
    // {8, 6, 4, 4, 1},
    // {1, 3, 6, 7, 9},
}
*/

func Day2PartOne() {
    count := 0

    reports, err := processProblemInput()
    if err != nil {
        log.Fatalf("Error %v", err)
        return 
    }

    for _, report := range reports {
        if isSafe(report) { count++ }
    }

    fmt.Println("Safe reports are:", count)
}

func Day2PartTwo() {
    count := 0

    reports, err := processProblemInput()
    if err != nil {
        log.Fatalf("Error %v", err)
        return 
    }

    for _, report := range reports {
        for i := range report {
            merged := append([]int{}, report[:i]...)
            merged = append(merged, report[i+1:]...)
            fmt.Println("merged -> ", merged)
            if isSafe(merged) { 
                count++ 
                break
            }
        }
    }

    fmt.Println("Safe reports are:", count)
}

func isSafe(report []int) bool {
    var isPositive bool 

    if (report[0] - report[1]) >= 0 {
        isPositive = true
    } else {
        isPositive = false
    }
    
    for i := range report[0:len(report)-1] { // from idx 1 to one before the last
        diff := report[i] - report[i+1]

        if diff < 0 && isPositive { return false }
        if diff == 0 { return false }
        if diff > 0 && !isPositive { return false }
        if helpers.Abs(diff) > 3 { return false }

    } 

    return true
}

// processProblemInput - processes this problem's data using the helper func
func processProblemInput() ([][]int, error) {
    // Day2PartOne called in main.go thus path = './day2/input.txt' 
    file, err := helpers.ReadInput("./day2/input.txt")
    if err != nil {
        return nil, err
    }

    // get the input and go through each digit on each line
    // and put it in the 2d array
    var reports [][]int

    for _, line := range file {
        var report []int

        // process each line to have a separate field per line
        fields := strings.Fields(line)
        for _, field := range fields {

            num, err := strconv.Atoi(field)
            if err != nil {
                return nil, err
            }

            report = append(report, num)
        }

        reports = append(reports, report)
    }

    return reports, nil
}
