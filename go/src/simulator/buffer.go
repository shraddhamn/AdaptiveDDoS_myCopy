package main

import (
	"math"
	"time"
	"fmt"
	
)

func enqueuePacket(pkt packet) {

	LOCK_INGRESS_CAP[pkt.ingress].Lock()
	if (INGRESS_CAP[pkt.ingress].availableBuffSpace - pkt.packet_len) > 0 {

		// LOCK_RECEIVE_COUNTER[pkt.ingress].Lock()
		// RECEIVE_COUNTER[pkt.ingress] +=1
		// LOCK_RECEIVE_COUNTER[pkt.ingress].Unlock()

		INGRESS_CAP[pkt.ingress].availableBuffSpace -= pkt.packet_len
		// _DEBUG.Printf("Function: enqueuePacket - Packet Added to Queue, Available Buffer space at %d = %f", pkt.ingress, INGRESS_CAP[pkt.ingress].availableBuffSpace)

		LOCK_INGRESS_CAP[pkt.ingress].Unlock()
		diagnose(pkt)
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

    if BACKLOG==256{
        fmt.Println("Backlog Full")  
    }
    
    s:=time.Now()
    s.Add(time.Second*5)
    backlog[BACKLOG]=s.String()
    BACKLOG+=1
    
    st:=time.Now()
    for i:=0; i< BACKLOG; i++{
    
        times=st.String()        
        if (times==backlog[BACKLOG]){
        
            for j:=i; j<BACKLOG-1; j++{
                backlog[j]=backlog[j+1]
            }
            
            BACKLOG-=1
        }
   
     }
    

	for i := 0; i < CONFIGURATION.INGRESS_LOC; i++ {
		// # LOCK_RECEIVE_COUNTER[i].Lock()

		// if(RECEIVE_COUNTER[i] > 0) {

		LOCK_INGRESS_CAP[i].Lock()
		pktsToDequeue := int(math.Ceil((INGRESS_CAP[i].vmQueue - INGRESS_CAP[i].availableBuffSpace) / PKT_LEN))
		if pktsToDequeue > INGRESS_CAP[i].numOfDequeuePkts {
			pktsToDequeue = INGRESS_CAP[i].numOfDequeuePkts
		}
		// if(float64(pktsToDequeue)*PKT_LEN > (INGRESS_CAP[i].cap - INGRESS_CAP[i].availableBuffSpace)) {

		// }
		// _DEBUG.Printf("Function: processPacket - %d packets to be processed, Available Buffer space at %d = %f", pktsToDequeue, i, INGRESS_CAP[i].availableBuffSpace)

		INGRESS_CAP[i].availableBuffSpace += (PKT_LEN * float64(pktsToDequeue))

		// _DEBUG.Printf("Function: processPacket - %d packets processed, Available Buffer space at %d = %f", pktsToDequeue, i, INGRESS_CAP[i].availableBuffSpace)
		LOCK_INGRESS_CAP[i].Unlock()
		// go dequeuePackets(pktsToDequeue,i)

		// _DEBUG.Printf("Function: processPacket - after diagnose pkts to dequeue = %d", pktsToDequeue)

		// }

		// # LOCK_RECEIVE_COUNTER[i].Unlock()

	}
}
