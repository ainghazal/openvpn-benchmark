package main

import (
	"bytes"
	"fmt"
	"os"
)

// tcNetem runs a tc netem command
func tcNetem(args ...string) {
	cmd := capCommand("tc", args...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(string(out))
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Stderr: %s\n", stderr.String())
		os.Exit(1)
	}
}

func setupLoss(config *config) {
	lossval := fmt.Sprintf("%d%%", config.loss)
	tcNetem("qdisc", "add", "dev", config.iface, "root", "netem", "loss", lossval)
}

func cleanupNetem(config *config) {
	tcNetem("qdisc", "del", "dev", config.iface, "root", "netem")
}
