package main

import (
	"math"
	"math/rand"
)

// func bohatei()
func changeCapacity(i int, newCap float64) {
	NUM_VMs[i] = int(math.Floor(newCap * 1.0 / CONFIGURATION.VM_COMPUTE_CAP))
	if NUM_VMs[i] == 0 {
		NUM_VMs[i] = 1
	}
	LOCK_INGRESS_CAP[i].Lock()
	// oldCap := INGRESS_CAP[i].cap
	INGRESS_CAP[i].cap = float64(NUM_VMs[i]) * CONFIGURATION.VM_COMPUTE_CAP
	/* snavalgu */
	//INGRESS_CAP[i].numOfDequeuePkts = int(INGRESS_CAP[i].cap / PKT_LEN)
	INGRESS_CAP[i].numOfDequeueBits = int(INGRESS_CAP[i].cap)
	oldQueue := INGRESS_CAP[i].vmQueue
	INGRESS_CAP[i].vmQueue = float64(CONFIGURATION.NUM_NIC_VM) * float64(NUM_VMs[i]) * CONFIGURATION.BUFF_SIZE
	INGRESS_CAP[i].availableBuffSpace = INGRESS_CAP[i].vmQueue - (oldQueue - INGRESS_CAP[i].availableBuffSpace)
	if INGRESS_CAP[i].availableBuffSpace < 0 {
		INGRESS_CAP[i].availableBuffSpace = oldQueue - INGRESS_CAP[i].availableBuffSpace
	}
	_DEBUG.Printf("Function: change Capacity, capacity at ingress %d = %f", i, INGRESS_CAP[i].cap)
	LOCK_INGRESS_CAP[i].Unlock()

}

func dynamicMitigation() {
	var newCap []float64
	newCap = make([]float64, CONFIGURATION.INGRESS_LOC)
	sum := 0.0
	for i := 0; i < CONFIGURATION.INGRESS_LOC; i++ {

		// _DEBUG.Printf("Function: dynamicMitigation, peak_traffic at ingress %d = %f", i, PEAK_TRAFFIC[i])
		// _DEBUG.Printf("Function: dynamicMitigation, min_traffic at ingress %d = %f", i, MIN_TRAFFIC[i])
		LOCK_CURR_TRAFFIC_STATS[i].Lock()
		perSec := float64(CURR_TRAFFIC_STATS[i]["total"]) / CONFIGURATION.EPOCH_TIME
		if perSec > PEAK_TRAFFIC[i] {
			PEAK_TRAFFIC[i] = perSec
		}

		if perSec < MIN_TRAFFIC[i] {
			MIN_TRAFFIC[i] = perSec
		}

		LOCK_CURR_TRAFFIC_STATS[i].Unlock()
		// # INGRESS_CAP[i] = random.uniform(MIN_TRAFFIC[i],PEAK_TRAFFIC[i])

		val := (rand.Float64()*(PEAK_TRAFFIC[i]-MIN_TRAFFIC[i]) + MIN_TRAFFIC[i]) * PKT_LEN
		newCap[i] = val
		sum += val
		// _DEBUG.Printf("Function: dynamicMitigation, peak_traffic at ingress %d = %f", i, PEAK_TRAFFIC[i]*PKT_LEN)
		// _DEBUG.Printf("Function: dynamicMitigation, min_traffic at ingress %d = %f", i, MIN_TRAFFIC[i]*PKT_LEN)
		// _DEBUG.Printf("Function: dynamicMitigation, val at ingress %d = %f", i, val)
		// changeCapacity(i, newCap)

	}
	for i := 0; i < CONFIGURATION.INGRESS_LOC; i++ {
		cap := (newCap[i] / sum) * CONFIGURATION.ISP_CAP
		// _DEBUG.Printf("Function: dynamicMitigation, capacity at ingress %d = %f, newcap[i] is %f, sum is %f", i, cap, newCap[i], sum)
		changeCapacity(i, cap)
	}
}

func staticMitigation() {

	for i := 0; i < CONFIGURATION.INGRESS_LOC; i++ {

		LOCK_CURR_TRAFFIC_STATS[i].Lock()

		AVG_TRAFFIC[i] = (AVG_TRAFFIC[i]*float64(WINDOW_COUNTER) + float64(CURR_TRAFFIC_STATS[i]["total"])) / (float64(WINDOW_COUNTER) + 1.0)
		perSec := (AVG_TRAFFIC[i] / CONFIGURATION.EPOCH_TIME) * PKT_LEN
		// if (float64(CURR_TRAFFIC_STATS[i]["total"]))*PKT_LEN > PEAK_TRAFFIC[i] {
		// 	PEAK_TRAFFIC[i] = float64(CURR_TRAFFIC_STATS[i]["total"]) * PKT_LEN
		// }

		// if float64(CURR_TRAFFIC_STATS[i]["total"])*PKT_LEN < MIN_TRAFFIC[i] {
		// 	MIN_TRAFFIC[i] = float64(CURR_TRAFFIC_STATS[i]["total"]) * PKT_LEN
		// }

		LOCK_CURR_TRAFFIC_STATS[i].Unlock()
		// # INGRESS_CAP[i] = random.uniform(MIN_TRAFFIC[i],PEAK_TRAFFIC[i])
		changeCapacity(i, perSec)

	}
}

func adaptiveMitigation() {

	for i := 0; i < CONFIGURATION.INGRESS_LOC; i++ {

		LOCK_CURR_TRAFFIC_STATS[i].Lock()

		AVG_TRAFFIC[i] = (AVG_TRAFFIC[i]*float64(WINDOW_COUNTER) + float64(CURR_TRAFFIC_STATS[i]["total"])) / (float64(WINDOW_COUNTER) + 1.0)
		perSec := (AVG_TRAFFIC[i] / CONFIGURATION.EPOCH_TIME) * PKT_LEN
		// if (float64(CURR_TRAFFIC_STATS[i]["total"]))*PKT_LEN > PEAK_TRAFFIC[i] {
		// 	PEAK_TRAFFIC[i] = float64(CURR_TRAFFIC_STATS[i]["total"]) * PKT_LEN
		// }

		// if float64(CURR_TRAFFIC_STATS[i]["total"])*PKT_LEN < MIN_TRAFFIC[i] {
		// 	MIN_TRAFFIC[i] = float64(CURR_TRAFFIC_STATS[i]["total"]) * PKT_LEN
		// }

		LOCK_CURR_TRAFFIC_STATS[i].Unlock()
		// # INGRESS_CAP[i] = random.uniform(MIN_TRAFFIC[i],PEAK_TRAFFIC[i])
		changeCapacity(i, perSec)

	}
}
