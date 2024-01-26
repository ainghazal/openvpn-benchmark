# 2024-01-26

I do another iteration for the side-to-side comparison, but with a dedicated interface (hardware NIC).

Setup 20% loss at each side, 100ms +- 50ms delay.

```
tc qdisc add dev enp2s0 root netem loss 20% delay 100ms 50ms

❯ ping vpn.local
PING vpn.local (192.168.8.200) 56(84) bytes of data.
64 bytes from vpn.local (192.168.8.200): icmp_seq=1 ttl=64 time=297 ms
64 bytes from vpn.local (192.168.8.200): icmp_seq=2 ttl=64 time=147 ms
64 bytes from vpn.local (192.168.8.200): icmp_seq=3 ttl=64 time=261 ms
64 bytes from vpn.local (192.168.8.200): icmp_seq=6 ttl=64 time=123 ms
64 bytes from vpn.local (192.168.8.200): icmp_seq=7 ttl=64 time=222 ms
64 bytes from vpn.local (192.168.8.200): icmp_seq=8 ttl=64 time=162 ms
64 bytes from vpn.local (192.168.8.200): icmp_seq=9 ttl=64 time=240 ms
64 bytes from vpn.local (192.168.8.200): icmp_seq=10 ttl=64 time=194 ms
64 bytes from vpn.local (192.168.8.200): icmp_seq=11 ttl=64 time=196 ms
64 bytes from vpn.local (192.168.8.200): icmp_seq=13 ttl=64 time=205 ms
64 bytes from vpn.local (192.168.8.200): icmp_seq=14 ttl=64 time=128 ms
64 bytes from vpn.local (192.168.8.200): icmp_seq=15 ttl=64 time=186 ms
64 bytes from vpn.local (192.168.8.200): icmp_seq=16 ttl=64 time=219 ms
64 bytes from vpn.local (192.168.8.200): icmp_seq=18 ttl=64 time=143 ms
^C
--- vpn.local ping statistics ---
19 packets transmitted, 14 received, 26.3158% packet loss, time 18053ms
rtt min/avg/max/mdev = 123.030/194.480/296.505/49.269 ms
```

# run at 20% loss (each way) + 100ms 50ms delay

![scatterplot for comparison](https://github.com/ainghazal/openvpn-benchmark/blob/main/scatterplot-openvpn-minivpn-20loss-100_50ms.png)


# openvpn failures

openvpn seems to choke at some point with receiving unexpected packets for the TLS handshake:
(then unrecoverable stall).

I *guess* this improves while using --tls-auth.

```
2024-01-26 18:17:13 TCP/UDP: Preserving recently used remote address: [AF_INET]192.168.8.200:1194
2024-01-26 18:17:13 Socket Buffers: R=[212992->212992] S=[212992->212992]
2024-01-26 18:17:13 UDPv4 link local: (not bound)
2024-01-26 18:17:13 UDPv4 link remote: [AF_INET]192.168.8.200:1194
2024-01-26 18:17:19 TLS Error: Unroutable control packet received from [AF_INET]192.168.8.200:1194 (si=3 op=P_ACK_V1)
2024-01-26 18:17:43 TLS Error: Unroutable control packet received from [AF_INET]192.168.8.200:1194 (si=3 op=P_ACK_V1)
2024-01-26 18:18:13 TLS Error: TLS key negotiation failed to occur within 60 seconds (check your network connectivity)
2024-01-26 18:18:13 TLS Error: TLS handshake failed
2024-01-26 18:18:13 SIGUSR1[soft,tls-error] received, process restarting
```

