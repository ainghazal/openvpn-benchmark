run-openvpn:
	 openvpn --disable-dco --management 127.0.0.1 6061 --management-hold --route-noexec --config data/client.ovpn
run-management:
	./openvpn-benchmark
setcap:
	sudo setcap 'cap_net_admin=ep' ./openvpn-benchmark

