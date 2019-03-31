package main

import (
	"fmt"
	"math"
	"../helper/go-cache"
	
)

//add Queue here for Customer


func addToBacklog(pkt packet) {

	if CONN_CUST == 256 {
		fmt.Println("Backlog Full")
		dropPacket(pkt)
	} else {
	
	Backlog_Cust[pkt.ingress].Add(pkt.src,1,cache.DefaultExpiration)
	CONN_CUST+=1
	fmt.Println(Backlog_Cust.itemCount())
    }
 }
	
	
	
func RemoveFromBacklog(pkt packet) {
    Backlog_Cust[pkt.ingress].Delete(pkt.src)
    CONN_CUST-=1
    fmt.Println(Backlog_Cust.itemCount())
}
	
	

func processPacket(attackType string) {
        
        //We will need to change the bitsToDequeue function for this or implement a Queue here
        bitsToDequeue := int(math.Ceil((INGRESS_CAP[i][attackType].vmQueue - INGRESS_CAP[i][attackType].availableBuffSpace)))        
        var pkt packet = dequeue(i)

        //checker function to check whether the attack packet is a dummy sent by attacker
        
        if pkt.checker == 1 {          
            send(dst,src,pkt)}
        
        
		if pkt.synFlag == 1 {
			addToBacklog(pkt)
		} else if pkt.ackFlag == 1 {
    		RemoveFromBacklog(pkt)
		}
        
}        

#	for i := 0; i < CONFIGURATION.INGRESS_LOC; i++ {
#
#		LOCK_INGRESS_CAP[i].Lock()
#
#		//pktsToDequeue := int(math.Ceil((INGRESS_CAP[i].vmQueue - INGRESS_CAP[i].availableBuffSpace) / PKT_LEN))
#		bitsToDequeue := int(math.Ceil((INGRESS_CAP[i][attackType].vmQueue - INGRESS_CAP[i][attackType].availableBuffSpace)))
#
#		for bitsToDequeue >= 0 {
#			var pkt packet = dequeue(i)
#
#			if pkt.synFlag == 1 {
#				addToBacklog(pkt)
#			} else if pkt.ackFlag == 1 {
#    			RemoveFromBacklog(pkt)
#			}
#			go diagnose(pkt)
#			bitsToDequeue -= int(pkt.packet_len)
#		}
#
#		INGRESS_CAP[i][attackType].availableBuffSpace += (float64(bitsToDequeue))
#
#		LOCK_INGRESS_CAP[i].Unlock()
#
#	}
}
