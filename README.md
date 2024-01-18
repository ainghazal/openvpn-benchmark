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


# reference

* [TUN/TAP device driver](https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/Documentation/networking/tuntap.rst)
