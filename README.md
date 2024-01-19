# openvpn-benchmark

micro-benchmarks and correctness tests for the openvpn implementation

# build

```
git clone https://github.com/ainghazal/gopenvpn ../
cat <<EOF > go.work
go 1.20

use (
    .
    ../gopenvpn
)
EOF
go get ./...
go build
sudo setcap 'cap_net_admin=ep' ./openvpn-benchmark
```

# run

```
./openvpn-benchmark -h
Usage of ./openvpn-benchmark:
  -count int
        how many times to run the experiment (default 1)
  -disable-dco
        disable dco module, if loaded (ref only)
  -file string
        file where to append results (default "results.json")
  -flavor string
        what impl to run (minivpn|ref) (default "ref")
  -iface string
        interface where to emulate network conditions (default "eth0")
  -loss int
        setup a specific packet loss % on the interface
```

# reference

* [TUN/TAP device driver](https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/Documentation/networking/tuntap.rst)
* [tc netem](https://www.man7.org/linux/man-pages/man8/tc-netem.8.html)
