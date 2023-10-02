package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLinesFromFile(filename string) ([]string, error) {
	// Open the text file for reading
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a slice to store the lines
	var lines []string

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read lines and append them to the slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {
	filename := "adventofcode.com_2022_day_1_input.txt"

	lines, err := readLinesFromFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the lines
	var curCalTotal = 0
	var elfCalTotals = []int{0}
	var elfIndex = 0

	for _, line := range lines {
		//fmt.Println(line)

		if line != "" {
			var calories = 0
			calories, err = strconv.Atoi(line)
			curCalTotal += calories
			elfCalTotals[elfIndex] = curCalTotal
		} else {
			elfIndex++
			elfCalTotals = append(elfCalTotals, 0)				
			curCalTotal = 0
		}
	}

	fmt.Println("There are this many elves:", len(elfCalTotals))
	var max = 0
	var elfIndexWithMax = 0
	for i := 0; i < len(elfCalTotals); i++ {
		fmt.Println(elfCalTotals[i])
		if elfCalTotals[i] > max {
			max = elfCalTotals[i]
			elfIndexWithMax = i
		}
	}

	fmt.Println("Elf with most calories is ", elfIndexWithMax, " and they have ", max, " calories")

}
