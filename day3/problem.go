package day3

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func PartOne() {

	file, err := os.ReadFile("./day3/input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	pattern := `mul\(\d{1,3},\d{1,3}\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(string(file), -1)

	totScore := calculate(matches)

	fmt.Println("tot score:", totScore)
}

func PartTwoAlt() {

	file, err := os.ReadFile("./day3/input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	ignorePattern := `don't\(\)(.*?)do\(\)`
	re := regexp.MustCompile(ignorePattern)

	// First replace all don't()...do() sections with just do()
	cleanText := re.ReplaceAllString(string(file), "do()")

	pattern := `mul\(\d{1,3},\d{1,3}\)`
	re = regexp.MustCompile(pattern)
	matches := re.FindAllString(cleanText, -1)

	totScore := calculate(matches)
	fmt.Println("Tot Score: ", totScore)

}

func calculate(matches []string) int {
	var totScore int

	for _, match := range matches {

		re := regexp.MustCompile(`\d+`)
		numbers := re.FindAllString(match, -1)

		var nums []int

		for _, num := range numbers {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}

		totScore += nums[0] * nums[1]
		// fmt.Println("numbers-> ", numbers)
	}

	return totScore
}
