package helpers

import (
	"bufio"
	"os"
)

// Returns absolute value
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Scan text file
func ReadInput(path string) ([]string, error) {
	file, err := os.Open(path)
    if err != nil {
        return nil, err
    }

    defer file.Close()

    var lines []string
    scan := bufio.NewScanner(file)

    for scan.Scan() {
        line := scan.Text()
        lines = append(lines, line)
    }

    return lines, scan.Err()
}
