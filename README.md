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

# simulating packet loss

TODO

```
sudo tc qdisc add dev enp38s0 root netem loss 10%
sudo tc qdisc show dev enp38s0
sudo tc qdisc del dev enp38s0 root netem
```


# reference

* [TUN/TAP device driver](https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/Documentation/networking/tuntap.rst)
* [tc netem](https://www.man7.org/linux/man-pages/man8/tc-netem.8.html)
