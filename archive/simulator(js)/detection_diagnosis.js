



// # let's fix 90% accuracy for each attack for now


function detect_UDP_Flood(udp_pkt) {
	prob = random.random(0.0,1)
	if(prob > UDP_DETEC_ACCURACY):
		# attack packet received
		return True

	return False
}




// function detect_TCP_SYN_Flood(tcp_pkt) {
// 	prob = random.random(0.0,1)
// 	if(prob > 0.9) {
// 		// # attack packet received
// 		return True
// 	}
// 	return False
// }



// function detect_NTP(udp_pkt) {
// 	prob = random.random(0.0,1)
// 	if(prob > 0.9) {
// 		// # attack packet received
// 		return True
// 	}
// 	return False
// }



// function detect_DNS_amp(udp_pkt) {
// 	prob = random.random(0.0,1)
// 	if(prob > 0.9) {
// 		// # attack packet received
// 		return True
// 	}
// 	return False
// }



function diagnose_UDP_Flood(udp_pkt) {

	UDP_Flood_traffic = 0

	if(detect_UDP_Flood(udp_pkt)) {
		UDP_Flood_traffic += udp_pkt.size
	}

	return UDP_Flood_traffic
}


// function diagnose_TCP_SYN_Flood(tcp_pkt) {

// 	TCP_SYN_Flood_traffic = 0

// 	if(detect_TCP_SYN_Flood(udp_pkt)){
// 		TCP_SYN_Flood_traffic += tcp_pkt.size
// 	}

// 	return TCP_SYN_Flood_traffic

// }


// function isUDP(pkt) {
// 	if(pkt.type == "UDP")
// }

function diagnose_tarffic(pkt) {

	if(isUDP(pkt)){
		attack_traffic_udp = diagnose_UDP_Flood(pkt)
		// attack_traffic_dns = diagnose_DNS_amp(pkt)
		// attack_traffic_ntp = diagnose_NTP(pkt)
	}
	if(isTCP(pkt)) {
		attack_traffic_tcp = diagnose_TCP_SYN_Flood(pkt)
	}


	traffic_stats["udp_flood"] = attack_traffic_udp
	traffic_stats["dns_amp"] = attack_traffic_dns
	traffic_stats["ntp"] = attack_traffic_ntp
	traffic_stats["tcp_syn"] = attack_traffic_tcp

	return diagnose_tarffic
}





