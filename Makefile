run-openvpn:
	 openvpn --disable-dco --management 127.0.0.1 6061 --management-hold --route-noexec --config data/client.ovpn
run-management:
	./openvpn-benchmark
setcap:
	sudo setcap 'cap_net_admin=ep' ./openvpn-benchmark
plot-separated:
	python parse.py experiment-openvpn-ref.json experiment-minivpn-ref.json | jq > comparison.json
	R -f plotData.R
plot-combined:
	python parse.py experiment-openvpn-2024-01-26.json | jq > comparison.json
	R -f plotData.R

