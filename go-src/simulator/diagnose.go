package main

import (
	"math/rand"
)

func detect_UDP_Flood(udp_pkt packet) bool {
	rnd := rand.Float64()
	if rnd < UDP_DETECT_ACCURACY {
		// _DEBUG.Printf("Function: detect_UDP_Flood - attack pkt detected")
		// # attack packet received
		return true
	}

	return false
}

func detect_TCP_SYN_Flood(tcp_pkt packet) bool {
	rnd := rand.Float64()
	if rnd < TCP_SYN_DETECT_ACCURACY {
		// # attack packet received
		return true
	}
	return false
}

func diagnose_UDP_Flood(pkt packet) {

	if detect_UDP_Flood(pkt) {
		LOCK_CURR_TRAFFIC_STATS[pkt.ingress].Lock()
		CURR_TRAFFIC_STATS[pkt.ingress]["UDP_FLOOD"] += 1
		CURR_TRAFFIC_STATS[pkt.ingress]["total"] += 1
		LOCK_CURR_TRAFFIC_STATS[pkt.ingress].Unlock()
	} else {
		LOCK_CURR_TRAFFIC_STATS[pkt.ingress].Lock()
		CURR_TRAFFIC_STATS[pkt.ingress]["total"] += 1
		LOCK_CURR_TRAFFIC_STATS[pkt.ingress].Unlock()
	}
}

func diagnose_TCP_SYN_Flood(pkt packet) {

	if detect_TCP_SYN_Flood(pkt) {
		LOCK_CURR_TRAFFIC_STATS[pkt.ingress].Lock()
		CURR_TRAFFIC_STATS[pkt.ingress]["TCP_SYN"] += 1
		CURR_TRAFFIC_STATS[pkt.ingress]["total"] += 1
		LOCK_CURR_TRAFFIC_STATS[pkt.ingress].Unlock()

	} else {
		LOCK_CURR_TRAFFIC_STATS[pkt.ingress].Lock()
		CURR_TRAFFIC_STATS[pkt.ingress]["total"] += 1
		LOCK_CURR_TRAFFIC_STATS[pkt.ingress].Unlock()
	}
}

func isUDP(pkt packet) bool {
	if pkt.protocol == "udp" {
		return true
	}
	return false
}

func isTCP(pkt packet) bool {
	if pkt.protocol == "tcp" {
		return true
	}
	return false
}

func diagnoseTraffic(pkt packet) {

	if pkt.attack_flag == true {
		if isUDP(pkt) {
			diagnose_UDP_Flood(pkt)
		}

		if isTCP(pkt) {
			diagnose_TCP_SYN_Flood(pkt)
		}
	} else {
		LOCK_CURR_TRAFFIC_STATS[pkt.ingress].Lock()
		CURR_TRAFFIC_STATS[pkt.ingress]["total"] += 1
		LOCK_CURR_TRAFFIC_STATS[pkt.ingress].Unlock()
	}
}
