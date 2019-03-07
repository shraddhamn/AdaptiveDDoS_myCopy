package main

import (
	"math"
)

func enqueuePacket(pkt packet) {

	LOCK_INGRESS_CAP[pkt.ingress].Lock()
	if (INGRESS_CAP[pkt.ingress].availableBuffSpace - pkt.packet_len) > 0 {

		// LOCK_RECEIVE_COUNTER[pkt.ingress].Lock()
		// RECEIVE_COUNTER[pkt.ingress] +=1
		// LOCK_RECEIVE_COUNTER[pkt.ingress].Unlock()

		INGRESS_CAP[pkt.ingress].availableBuffSpace -= pkt.packet_len
		// _DEBUG.Printf("Function: enqueuePacket - Packet Added to Queue, Available Buffer space at %d = %f", pkt.ingress, INGRESS_CAP[pkt.ingress].availableBuffSpace)
		enqueue(pkt, pkt.ingress)
		LOCK_INGRESS_CAP[pkt.ingress].Unlock()
		//diagnose(pkt)
		// BUFFER[pkt.ingress].Add(pkt)
		// BUFFER[pkt.ingress] <- pkt   // Send v to channel ch

	} else {
		LOCK_INGRESS_CAP[pkt.ingress].Unlock()
		dropPacket(pkt)
	}

}

func dropPacket(pkt packet) {

	if pkt.attack_flag == false {

		LOCK_legitimateDropCounter[pkt.ingress].Lock()
		legitimateDropCounter[pkt.ingress] += 1
		// _DEBUG.Printf("Function: dropPacket - Legitimate packet dropped, legitimateDropCounter = %d",legitimateDropCounter[pkt.ingress])
		LOCK_legitimateDropCounter[pkt.ingress].Unlock()

		//  	} else {
		// 	attackDropCounter[pkt.ingress] +=1
		//      // # _DEBUG.Printf(f"Function: dropPacket - Attack packet dropped, attackDropCounter = {attackDropCounter[pkt.ingress]}")
	}
}

// func dequeuePackets(pktsToDequeue int, ingress int) {
// 	for j := 0 ; j < pktsToDequeue ; j++ {
// 				// pkt := <-BUFFER[i]
// 		pkt,ok  := (BUFFER[ingress].Next()).(packet)
// 		if(ok) {
// 			// enqueuePacket(pkt)
// 			diagnose(pkt)
// 		}
// 	}
// }
// func processing() {
// 	for {
// 		processPacket()
// 		// time.Sleep(10 * time.Microsecond)
// 	}
// }

func processPacket() {

	for i := 0; i < CONFIGURATION.INGRESS_LOC; i++ {
		// # LOCK_RECEIVE_COUNTER[i].Lock()

		// if(RECEIVE_COUNTER[i] > 0) {

		LOCK_INGRESS_CAP[i].Lock()

		//pktsToDequeue := int(math.Ceil((INGRESS_CAP[i].vmQueue - INGRESS_CAP[i].availableBuffSpace) / PKT_LEN))
		bitsToDequeue := int(math.Ceil((INGRESS_CAP[i].vmQueue - INGRESS_CAP[i].availableBuffSpace)))
		// if pktsToDequeue > INGRESS_CAP[i].numOfDequeuePkts {
		// 	pktsToDequeue = INGRESS_CAP[i].numOfDequeuePkts
		// }
		for bitsToDequeue >= 0 {
			var pkt packet = dequeue(i)
			bitsToDequeue -= int(pkt.packet_len)
		}

		// if(float64(pktsToDequeue)*PKT_LEN > (INGRESS_CAP[i].cap - INGRESS_CAP[i].availableBuffSpace)) {

		// }
		// _DEBUG.Printf("Function: processPacket - %d packets to be processed, Available Buffer space at %d = %f", pktsToDequeue, i, INGRESS_CAP[i].availableBuffSpace)

		//INGRESS_CAP[i].availableBuffSpace += (PKT_LEN * float64(pktsToDequeue))

		INGRESS_CAP[i].availableBuffSpace += (float64(bitsToDequeue))

		// _DEBUG.Printf("Function: processPacket - %d packets processed, Available Buffer space at %d = %f", pktsToDequeue, i, INGRESS_CAP[i].availableBuffSpace)
		LOCK_INGRESS_CAP[i].Unlock()
		// go dequeuePackets(pktsToDequeue,i)

		// _DEBUG.Printf("Function: processPacket - after diagnose pkts to dequeue = %d", pktsToDequeue)

		// }

		// # LOCK_RECEIVE_COUNTER[i].Unlock()

	}
}
