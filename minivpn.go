package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

// runMiniVPN runs the minivpn implementation
func runMiniVPN(config *config) {
	var wg sync.WaitGroup

	// wait a bit. we're waiting for openvpn in a different moment.
	time.Sleep(time.Second)

	start := time.Now()

	wg.Add(1)

	// launch minivpn
	go func() {
		defer wg.Done()

		cmd := capCommand("./minivpn", "data/client.ovpn")
		var stderr bytes.Buffer
		cmd.Stderr = &stderr

		// Create a pipe to capture the command's stdout
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println("Error creating StdoutPipe:", err)
			return
		}

		// Start the command
		log.Println("start")
		if err := cmd.Start(); err != nil {
			fmt.Println("Error starting command:", err)
			return
		}

		// Create a scanner to read the stdout in real-time
		scanner := bufio.NewScanner(stdout)

		// Start a goroutine to handle the output
		go func() {
			for scanner.Scan() {
				line := scanner.Text()

				// Process the log line as needed
				if strings.HasPrefix(line, "Local") {
					fmt.Println(line)
					continue
				}
				if strings.HasPrefix(line, "Gateway") {
					fmt.Println(line)
					fmt.Println("done!")
					if config.debug {
						fmt.Printf("Stderr: %s\n", stderr.String())
					}

					wg.Done()
					return
				}
			}
		}()

		// Wait for the command to finish
		if err := cmd.Wait(); err != nil {
			fmt.Println("Command finished with error:", err)
		}
		// Close the stdout pipe
		stdout.Close()
		// Print the stderr buffer
		fmt.Printf("Stderr: %s\n", stderr.String())
	}()

	fmt.Println("waitttt")

	wg.Wait()

	elapsed := time.Since(start)
	result := result{
		Time:    time.Now(),
		Loss:    config.loss,
		Flavor:  "minivpn",
		Elapsed: elapsed.String(),
	}
	r, _ := json.Marshal(result)
	fmt.Println(string(r))

	if config.file != "" {
		appendToFile(config.file, string(r)+",\n")
	}
}
