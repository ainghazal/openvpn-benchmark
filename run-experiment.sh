#!/bin/sh
EXPERIMENT=experiment-02.json
IFACE=enp38s0

./openvpn-benchmark -iface $IFACE -loss 2 -count 5 -flavor openvpn -file $EXPERIMENT
sleep 5
./openvpn-benchmark -iface $IFACE -loss 2 -count 5 -flavor minivpn -file $EXPERIMENT -debug
sleep 5

./openvpn-benchmark -iface $IFACE -loss 4 -count 5 -flavor openvpn -file $EXPERIMENT
sleep 5
./openvpn-benchmark -iface $IFACE -loss 4 -count 5 -flavor minivpn -file $EXPERIMENT -debug
sleep 5

./openvpn-benchmark -iface $IFACE -loss 6 -count 5 -flavor openvpn -file $EXPERIMENT
sleep 5
./openvpn-benchmark -iface $IFACE -loss 6 -count 5 -flavor minivpn -file $EXPERIMENT -debug
sleep 5

./openvpn-benchmark -iface $IFACE -loss 8 -count 5 -flavor openvpn -file $EXPERIMENT
sleep 5
./openvpn-benchmark -iface $IFACE -loss 8 -count 5 -flavor minivpn -file $EXPERIMENT -debug
sleep 5

./openvpn-benchmark -iface $IFACE -loss 10 -count 5 -flavor openvpn -file $EXPERIMENT
sleep 5
./openvpn-benchmark -iface $IFACE -loss 10 -count 5 -flavor minivpn -file $EXPERIMENT -debug
sleep 5

# ./openvpn-benchmark -iface $IFACE -loss 10 -count 5 -flavor openvpn -file $EXPERIMENT
# ./openvpn-benchmark -iface $IFACE -loss 10 -count 5 -flavor minivpn -file $EXPERIMENT
# 
# ./openvpn-benchmark -iface $IFACE -loss 12 -count 20 -flavor openvpn -file $EXPERIMENT
# ./openvpn-benchmark -iface $IFACE -loss 12 -count 20 -flavor minivpn -file $EXPERIMENT
# 
# ./openvpn-benchmark -iface $IFACE -loss 14 -count 20 -flavor openvpn -file $EXPERIMENT
# ./openvpn-benchmark -iface $IFACE -loss 14 -count 20 -flavor minivpn -file $EXPERIMENT
# 
# ./openvpn-benchmark -iface $IFACE -loss 16 -count 20 -flavor openvpn -file $EXPERIMENT
# ./openvpn-benchmark -iface $IFACE -loss 16 -count 20 -flavor minivpn -file $EXPERIMENT
# 
# ./openvpn-benchmark -iface $IFACE -loss 20 -count 20 -flavor openvpn -file $EXPERIMENT
# ./openvpn-benchmark -iface $IFACE -loss 20 -count 20 -flavor minivpn -file $EXPERIMENT
# 
# ./openvpn-benchmark -iface $IFACE -loss 24 -count 20 -flavor openvpn -file $EXPERIMENT
# ./openvpn-benchmark -iface $IFACE -loss 24 -count 20 -flavor minivpn -file $EXPERIMENT
# 
# ./openvpn-benchmark -iface $IFACE -loss 28 -count 20 -flavor openvpn -file $EXPERIMENT
# ./openvpn-benchmark -iface $IFACE -loss 28 -count 20 -flavor minivpn -file $EXPERIMENT
# 
# ./openvpn-benchmark -iface $IFACE -loss 32 -count 20 -flavor openvpn -file $EXPERIMENT
# ./openvpn-benchmark -iface $IFACE -loss 32 -count 20 -flavor minivpn -file $EXPERIMENT
# 
# ./openvpn-benchmark -iface $IFACE -loss 36 -count 20 -flavor openvpn -file $EXPERIMENT
# ./openvpn-benchmark -iface $IFACE -loss 36 -count 20 -flavor minivpn -file $EXPERIMENT
# 
# ./openvpn-benchmark -iface $IFACE -loss 40 -count 20 -flavor openvpn -file $EXPERIMENT
# ./openvpn-benchmark -iface $IFACE -loss 40 -count 20 -flavor minivpn -file $EXPERIMENT
