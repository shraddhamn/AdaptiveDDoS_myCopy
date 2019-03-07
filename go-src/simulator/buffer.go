package main

import (
	"fmt"
	"math"
	"time"
)

func enqueuePacket(pkt packet) {

	LOCK_INGRESS_CAP[pkt.ingress].Lock()
	if (INGRESS_CAP[pkt.ingress].availableBuffSpace - pkt.packet_len) > 0 {

		INGRESS_CAP[pkt.ingress].availableBuffSpace -= pkt.packet_len
		// _DEBUG.Printf("Function: enqueuePacket - Packet Added to Queue, Available Buffer space at %d = %f", pkt.ingress, INGRESS_CAP[pkt.ingress].availableBuffSpace)
		enqueue(pkt, pkt.ingress)
		LOCK_INGRESS_CAP[pkt.ingress].Unlock()

	} else {
		LOCK_INGRESS_CAP[pkt.ingress].Unlock()
		dropPacket(pkt)
	}

}

func dropPacket(pkt packet) {

	if pkt.attack_flag == false {

		LOCK_legitimateDropCounter[pkt.ingress].Lock()
		legitimateDropCounter[pkt.ingress] += 1
		_DEBUG.Printf("Function: dropPacket - Legitimate packet dropped, legitimateDropCounter = %d", legitimateDropCounter[pkt.ingress])
		LOCK_legitimateDropCounter[pkt.ingress].Unlock()

		//  	} else {
		// 	attackDropCounter[pkt.ingress] +=1
		//      // # _DEBUG.Printf(f"Function: dropPacket - Attack packet dropped, attackDropCounter = {attackDropCounter[pkt.ingress]}")
	}
}

func addToBacklog(pkt packet) {
	if CONN_IN_BACKLOG == 256 {
		fmt.Println("Backlog Full")
		dropPacket(pkt)
	} else {
		s := time.Now()
		s.Add(time.Second * 5)
		BACKLOG[CONN_IN_BACKLOG] = s.String()
		CONN_IN_BACKLOG += 1

		st := time.Now()
		for i := 0; i < CONN_IN_BACKLOG; i++ {

			times = st.String()
			if times == BACKLOG[CONN_IN_BACKLOG] {

				for j := i; j < CONN_IN_BACKLOG-1; j++ {
					BACKLOG[j] = BACKLOG[j+1]
				}

				CONN_IN_BACKLOG -= 1
			}

		}
	}
}

func processPacket() {

	for i := 0; i < CONFIGURATION.INGRESS_LOC; i++ {

		LOCK_INGRESS_CAP[i].Lock()

		//pktsToDequeue := int(math.Ceil((INGRESS_CAP[i].vmQueue - INGRESS_CAP[i].availableBuffSpace) / PKT_LEN))
		bitsToDequeue := int(math.Ceil((INGRESS_CAP[i].vmQueue - INGRESS_CAP[i].availableBuffSpace)))

		for bitsToDequeue >= 0 {
			var pkt packet = dequeue(i)
			if pkt.synFlag == 1 {
				addToBacklog(pkt)
			}
			go diagnose(pkt)
			bitsToDequeue -= int(pkt.packet_len)
		}

		INGRESS_CAP[i].availableBuffSpace += (float64(bitsToDequeue))

		LOCK_INGRESS_CAP[i].Unlock()

	}
}
