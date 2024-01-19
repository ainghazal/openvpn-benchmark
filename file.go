package main

import (
	"fmt"
	"os"
)

func appendToFile(filename, content string) error {
	// Open the file in append mode with write permissions
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the content to the file
	if _, err := file.WriteString(content); err != nil {
		return err
	}

	fmt.Printf("wrote result to %s\n", filename)
	return nil
}
