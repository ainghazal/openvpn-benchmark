//
// Benchmark bootstrapping time against a reference OpenVPN server implementation.
//

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"time"

	"golang.org/x/sys/unix"
)

type config struct {
	flavor     string
	disableDCO bool
	loss       int
	iface      string
	count      int
	file       string
	debug      bool
}

type result struct {
	Time    time.Time `json:"t"`
	Loss    int       `json:"loss"`
	Flavor  string    `json:"flavor"`
	Elapsed string    `json:"elapsed"`
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

func capCommandTimeout(command string, args ...string) (*exec.Cmd, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*TLS_TIMEOUT)
	cmd := exec.CommandContext(ctx, command, args...)
	cmd.SysProcAttr = &unix.SysProcAttr{
		AmbientCaps: []uintptr{unix.CAP_NET_ADMIN},
	}
	return cmd, cancel
}

func main() {

	config := &config{}

	flag.IntVar(&config.count, "count", 1, "how many times to run the experiment")
	flag.StringVar(&config.flavor, "flavor", "ref", "what impl to run (minivpn|ref)")
	flag.StringVar(&config.iface, "iface", "eth0", "interface where to emulate network conditions")
	flag.BoolVar(&config.disableDCO, "disable-dco", false, "disable dco module, if loaded (ref only)")
	flag.IntVar(&config.loss, "loss", 0, "setup a specific packet loss % on the interface")
	flag.StringVar(&config.file, "file", "results.json", "file where to append results")
	flag.BoolVar(&config.debug, "debug", false, "print logs")
	flag.Parse()

	log.Println("iface:", config.iface)
	log.Println("file:", config.file)

	maybeCleanupNetem(config)

	defer cleanupNetem(config)
	setupLoss(config)

	for i := 1; i < config.count+1; i++ {
		switch config.flavor {
		case "minivpn":
			fmt.Printf("yo, run minivpn (%d)!\n", i)
			runMiniVPN(config)
		default:
			fmt.Printf("bojack openvpn, run wild! go %d!\n", i)
			runReference(config)
		}
		fmt.Println("run done, waiting for things to settle...")
		time.Sleep(time.Second * 5)
	}
}
