#!/bin/sh
EXPERIMENT=experiment-openvpn-2024-01-26.json
IFACE=eno1

./openvpn-benchmark -iface $IFACE -loss 20 -count 10 -flavor openvpn -debug -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 20 -count 10 -flavor minivpn -debug -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 20 -count 10 -flavor openvpn -debug -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 20 -count 10 -flavor minivpn -debug -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 20 -count 10 -flavor openvpn -debug -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 20 -count 10 -flavor minivpn -debug -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 20 -count 10 -flavor openvpn -debug -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 20 -count 10 -flavor minivpn -debug -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 20 -count 10 -flavor openvpn -debug -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 20 -count 10 -flavor minivpn -debug -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 20 -count 10 -flavor openvpn -debug -file $EXPERIMENT
./openvpn-benchmark -iface $IFACE -loss 20 -count 10 -flavor minivpn -debug -file $EXPERIMENT
