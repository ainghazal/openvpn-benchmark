package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/NordSecurity/gopenvpn/openvpn"
	"golang.org/x/sys/unix"
)

func main() {
	var disableDCO bool

	flag.BoolVar(&disableDCO, "disable-dco", false, "disable dco module, if loaded")
	flag.Parse()

	var wg sync.WaitGroup

	vpnargs := []string{}
	if disableDCO {
		vpnargs = append(vpnargs, "--disable-dco")
	}
	vpnargs = append(vpnargs, []string{
		"--management", "127.0.0.1", "6061",
		"--management-hold", "--route-noexec",
		"--config", "data/client.ovpn"}...)

	go func() {
		defer wg.Done()
		wg.Add(1)

		cmd := exec.Command("openvpn", vpnargs...)

		cmd.SysProcAttr = &unix.SysProcAttr{
			AmbientCaps: []uintptr{unix.CAP_NET_ADMIN},
		}
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(string(out))
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(string(out))
	}()

	// let's give some time to openvpn to start and wait for us
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

	if err := client.HoldRelease(); err != nil {
		panic(err)
	}

	go func() {
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
					}
				default:
				}
			}
		}
	}()
	wg.Wait()
}
