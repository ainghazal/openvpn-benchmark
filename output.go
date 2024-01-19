package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func getNextSequentialFilename(folderPath, baseFilename string) (string, error) {
	// List files in the folder
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return "", err
	}

	// Create a regular expression to match the base filename with a sequential number
	pattern := fmt.Sprintf("^%s_(\\d+)\\..*$", baseFilename)
	re := regexp.MustCompile(pattern)

	// Extract existing sequential numbers
	var numbers []int
	for _, file := range files {
		match := re.FindStringSubmatch(file.Name())
		if match != nil {
			num, err := strconv.Atoi(match[1])
			if err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	// Sort existing numbers in ascending order
	sort.Ints(numbers)

	// Find the next sequential number
	nextNumber := 1
	if len(numbers) > 0 {
		nextNumber = numbers[len(numbers)-1] + 1
	}

	// Create the new filename
	newFilename := fmt.Sprintf("%s_%d.txt", baseFilename, nextNumber)
	return newFilename, nil
}
