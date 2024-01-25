#!/bin/sh
EXPERIMENT=experiment-openvpn-ref.json
IFACE=enp38s0

./openvpn-benchmark -iface $IFACE -loss 2 -count 5 -flavor openvpn -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 4 -count 5 -flavor openvpn -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 6 -count 5 -flavor openvpn -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 8 -count 5 -flavor openvpn -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 10 -count 5 -flavor openvpn -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 14 -count 5 -flavor openvpn -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 16 -count 5 -flavor openvpn -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 20 -count 5 -flavor openvpn -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 25 -count 5 -flavor openvpn -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 30 -count 5 -flavor openvpn -file $EXPERIMENT
