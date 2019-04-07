package main

import (
	"math"

	"../helper/fifo"
)

// var (
// 	prevDropCount []int
// 	prevReceiveCount []int
// )

var (
	availableOutgoingQueueSpace float64
)

func initializeISP() {

	_DEBUG.Printf("Function: initializeISP- Initializing TRAFFIC_STATS variables")

	PEAK_TRAFFIC = make([]float64, CONFIGURATION.INGRESS_LOC)
	MIN_TRAFFIC = make([]float64, CONFIGURATION.INGRESS_LOC)
	AVG_TRAFFIC = make([]float64, CONFIGURATION.INGRESS_LOC)
	// RECEIVE_COUNTER = make([]int,CONFIGURATION.INGRESS_LOC)
	legitimateDropCounter = make([]int, CONFIGURATION.INGRESS_LOC)

	// prevDropCount = make([]int,CONFIGURATION.INGRESS_LOC)
	// prevReceiveCount = make([]int,CONFIGURATION.INGRESS_LOC)

	// processCounter = make([]int,CONFIGURATION.INGRESS_LOC)
	// BUFFER = make ([]fifo.Queue, CONFIGURATION.INGRESS_LOC)
	// BUFFER = make([]*fifo.Queue, CONFIGURATION.INGRESS_LOC)

	for i := 0; i < CONFIGURATION.INGRESS_LOC; i++ {
		var m map[string]float64
		m = make(map[string]float64)
		m["total"] = 0
		m["UDP_FLOOD"] = 0
		m["TCP_SYN"] = 0
		m["DNS_AMP"] = 0
		CURR_TRAFFIC_STATS = append(CURR_TRAFFIC_STATS, m)
		MIN_TRAFFIC[i] = math.Inf(0)
		pktQueue = append(pktQueue, fifo.NewQueue())
		// Backlog_Queue = append(Backlog_Queue,cache.New(5*time.Second, 5*time.Second))
		// PREV_TRAFFIC_STATS = append(PREV_TRAFFIC_STATS,m)
		// BUFFER[i] = fifo.NewQueue()

	}

}

func countDroppedPackets() {

	for i := 0; i < CONFIGURATION.INGRESS_LOC; i++ {

		LOCK_legitimateDropCounter[i].Lock()
		dropped := legitimateDropCounter[i]
		legitimateDropCounter[i] = 0
		LOCK_legitimateDropCounter[i].Unlock()

		_DEBUG.Printf("Function: countDroppedPackets, dropped Packets at ingress %d = %d", i, dropped)
		_INFO.Printf("Dropped_Packets %d ingress %d", dropped, i)

		// fmt.Printf("%d",dropped)
	}
	// # loggings.error('This should go to both console and file')
}

func wastedResources(total []map[string]float64) {

	for i := 0; i < CONFIGURATION.INGRESS_LOC; i++ {

		// LOCK_RECEIVE_COUNTER[i].Lock()

		// _INFO.Printf("packets received per window at ingress %d = %d Mbps",i, receivedPktsPerWIndow)

		// prevReceiveCount[i] = total[i]["total"]
		// LOCK_RECEIVE_COUNTER[i].Unlock()

		LOCK_INGRESS_CAP[i].Lock()

		for _, element := range ATTACK_TYPES {
			LOCK_CURR_TRAFFIC_STATS[i].Lock()
			receivedBitsPerWIndow := total[i][element]
			LOCK_CURR_TRAFFIC_STATS[i].Unlock()
			wastedCap := INGRESS_CAP[i][element].cap - (float64(receivedBitsPerWIndow) / CONFIGURATION.EPOCH_TIME)

			// fmt.Printf("%f",wastedCap)

			// _DEBUG.Printf("Function: wastedResources, wasted resources at ingress %d = %v Mbps",i, wastedCap)
			_INFO.Printf("Wasted_resources_Mbps %v ingress %d attackType %s", wastedCap, i, element)
		}
		LOCK_INGRESS_CAP[i].Unlock()

		// #

		// # print wastedCap
	}
}

func collectStats() {

	_INFO.Printf("STATS FOR WINDOW %d - START", WINDOW_COUNTER)
	// var total []map[string]int
	// total = make([]map[string]int, CONFIGURATION.INGRESS_LOC)
	// dest = rand.Intn(CONFIGURATION.INGRESS_LOC)
	wastedResources(CURR_TRAFFIC_STATS)
	mitigate()
	for i := 0; i < CONFIGURATION.INGRESS_LOC; i++ {

		LOCK_CURR_TRAFFIC_STATS[i].Lock()
		// PREV_TRAFFIC_STATS[i]["total"] = CURR_TRAFFIC_STATS[i]["total"]
		_INFO.Printf("Total_Traffic %v Ingress %d", CURR_TRAFFIC_STATS[i]["total"], i)
		// total = append(total,CURR_TRAFFIC_STATS[i]["total"])
		// t := CURR_TRAFFIC_STATS[i]["total"]
		// copy(total,CURR_TRAFFIC_STATS)
		// wastedResources(t,i)
		CURR_TRAFFIC_STATS[i]["total"] = 0

		// PREV_TRAFFIC_STATS[i]["udp_flood"] = CURR_TRAFFIC_STATS[i]["udp_flood"]
		_INFO.Printf("Total_UDP_Flood %v Ingress %d", CURR_TRAFFIC_STATS[i]["udp_flood"], i)
		CURR_TRAFFIC_STATS[i]["udp_flood"] = 0

		// PREV_TRAFFIC_STATS[i]["tcp_syn"] = CURR_TRAFFIC_STATS[i]["tcp_syn"]
		_INFO.Printf("Total_TCP_Syn %v Ingress %d", CURR_TRAFFIC_STATS[i]["tcp_syn"], i)
		CURR_TRAFFIC_STATS[i]["tcp_syn"] = 0

		// PREV_TRAFFIC_STATS[i]["dns_amp"] = CURR_TRAFFIC_STATS[i]["dns_amp"]
		_INFO.Printf("Total_DNS_amp %v Ingress %d", CURR_TRAFFIC_STATS[i]["dns_amp"], i)
		CURR_TRAFFIC_STATS[i]["dns_amp"] = 0
		LOCK_CURR_TRAFFIC_STATS[i].Unlock()

	}
	WINDOW_COUNTER += 1

	countDroppedPackets()

	_INFO.Printf("STATS FOR WINDOW %d - END \n\n", WINDOW_COUNTER)
}

func enqueueOutgoingQueue(pkt packet) {
	if availableOutgoingQueueSpace-pkt.packet_len > 0 {
		outgoingQueue.Add(pkt)
	} else {
		// TODO ask from where to drop and what coiunter to increment
		//dropPacket(pkt)
		_DEBUG.Printf("Function: enqueueOutgoingQueue - outgoing packet dropped")

	}
}

func dequeueOutgoingQueue() packet {
	return outgoingQueue.Next().(packet)
}

// call
func transmit(pkt packet) {
	// find number of dequeue bits
	dequeueOutgoingQueue()
	enqueueTarget(pkt)
}
