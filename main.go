//
// Benchmark bootstrapping time against a reference OpenVPN server implementation.
//

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/NordSecurity/gopenvpn/openvpn"
	"golang.org/x/sys/unix"
)

type config struct {
	flavor     string
	disableDCO bool
	loss       int
	iface      string
}

// capCommand runs a command with a specific set of capabilities.
// the binary itself needs to previously have had these capabilities set.
func capCommand(command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.SysProcAttr = &unix.SysProcAttr{
		AmbientCaps: []uintptr{unix.CAP_NET_ADMIN},
	}
	return cmd
}

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

func main() {

	config := &config{}

	flag.StringVar(&config.flavor, "flavor", "ref", "what impl to run (minivpn|ref)")
	flag.StringVar(&config.iface, "iface", "eth0", "interface where to emulate network conditions")
	flag.BoolVar(&config.disableDCO, "disable-dco", false, "disable dco module, if loaded")
	flag.IntVar(&config.loss, "loss", 0, "setup a specific packet loss %% on the interface")
	flag.Parse()

	log.Println("iface:", config.iface)

	defer cleanupNetem(config)
	setupLoss(config)

	switch config.flavor {
	case "minivpn":
		fmt.Println("yo, run minivpn!")
		runMiniVPN(config)
	default:
		fmt.Println("bojack openvpn, run wild! go!")
		runReference(config)
	}
}

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
	fmt.Println("flavor: minivpn")
	fmt.Printf("loss: %d%%\n", config.loss)
	fmt.Println("elapsed:", elapsed)
}

// runReference will run the openvpn reference implementation
func runReference(config *config) {
	var wg sync.WaitGroup

	vpnargs := []string{}

	if config.disableDCO {
		vpnargs = append(vpnargs, "--disable-dco")
	}

	// launch openvpn on the background, and make it wait for the management interface to relase the hold
	go func() {
		defer wg.Done()
		wg.Add(1)

		vpnargs = append(vpnargs, []string{
			"--management", "127.0.0.1", "6061",
			"--management-hold", "--route-noexec",
			"--config", "data/client.ovpn"}...)

		log.Println("openvpn", vpnargs)

		cmd := capCommand("openvpn", vpnargs...)
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(string(out))
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(string(out))
	}()

	// let's give some time to openvpn to start and wait for us?
	time.Sleep(time.Second)

	eventCh := make(chan openvpn.Event, 10)

	client, err := openvpn.Dial("127.0.0.1:6061", eventCh)
	if err != nil {
		panic(err)
	}

	if err := client.SetEchoEvents(true); err != nil {
		panic(err)
	}
	if err := client.SetLogEvents(true); err != nil {
		panic(err)
	}
	log.Println("start")
	start := time.Now()

	if err := client.HoldRelease(); err != nil {
		panic(err)
	}

	t := make(chan time.Time, 1)

	go func(t chan time.Time) {
		defer wg.Done()
		wg.Add(1)
		for {
			select {
			case ev := <-eventCh:
				if ev == nil {
					return
				}
				repr := ev.String()
				switch strings.HasPrefix(repr, "LOG") {
				case true:
					logLine := ev.(*openvpn.LogEvent).Line()
					if strings.Contains(logLine, "Initialization Sequence Completed") {
						log.Println("tunnel established!")
						t <- time.Now()
						client.SendSignal("SIGTERM")
						return
					}
				default:
				}
			}
		}
	}(t)

	wg.Wait()

	finished := <-t
	elapsed := finished.Sub(start)
	fmt.Println("flavor: openvpn")
	fmt.Printf("loss: %d%%\n", config.loss)
	fmt.Println("elapsed:", elapsed)
}
