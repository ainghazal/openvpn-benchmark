package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/NordSecurity/gopenvpn/openvpn"
)

const (
	TLS_TIMEOUT = 61
)

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

		cmd, cancel := capCommandTimeout("openvpn", vpnargs...)
		defer cancel()
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(string(out))
			fmt.Println("Error:", err)
			fmt.Println("openvpn command exited")
			return
		}
		fmt.Println("openvpn command exited")
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

	go func(chan time.Time) {
		defer wg.Done()
		wg.Add(1)
		timeout := time.NewTimer(time.Second * TLS_TIMEOUT)

		for {
			select {
			case <-timeout.C:
				fmt.Println("tls timeout!")
				t <- time.Now()
				return

			case ev := <-eventCh:
				if ev == nil {
					t <- time.Now()
					return
				}
				repr := ev.String()
				fmt.Println(repr)
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
	fmt.Println("done waiting")

	finished := <-t
	elapsed := finished.Sub(start)

	fmt.Println("elapsed:", elapsed)

	result := result{
		Time:    time.Now(),
		Loss:    config.loss,
		Flavor:  "openvpn",
		Elapsed: elapsed.String(),
	}
	r, _ := json.Marshal(result)
	fmt.Println(string(r))

	if config.file != "" {
		appendToFile(config.file, string(r)+",\n")
	}
}
