package main

import (
	"bytes"
	"fmt"
	"os"
)

// tcNetem runs a tc netem command
func tcNetem(canFail bool, args ...string) {
	cmd := capCommand("tc", args...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if canFail {
		return
	}

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(string(out))
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Stderr: %s\n", stderr.String())
		os.Exit(1)
	}
}

func setupLoss(config *config) {
	fmt.Printf("Setting Loss to %v%%\n", config.loss)
	lossval := fmt.Sprintf("%v%%", config.loss)
	tcNetem(false, "qdisc", "add", "dev", config.iface, "root", "netem", "loss", lossval,
		"delay", "100ms")
}

func cleanupNetem(config *config) {
	tcNetem(false, "qdisc", "del", "dev", config.iface, "root", "netem")
}

func maybeCleanupNetem(config *config) {
	tcNetem(true, "qdisc", "del", "dev", config.iface, "root", "netem")
}
