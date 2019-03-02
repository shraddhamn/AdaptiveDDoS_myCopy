package main

import "math"

type VM struct {
	cap     float64
	vmQueue float64
	//numOfDequeuePkts   int
	availableBuffSpace float64
	numOfDequeueBits   int
}

func initializeDefense() {
	_DEBUG.Printf("Function: initializeAdaptive - Initialize inital capapcity of ingress locations")
	total_num_vms := CONFIGURATION.ISP_CAP / CONFIGURATION.VM_COMPUTE_CAP
	for i := 0; i < CONFIGURATION.INGRESS_LOC; i++ {
		NUM_VMs = append(NUM_VMs, int(math.Floor(total_num_vms*1.0/float64(CONFIGURATION.INGRESS_LOC))))
		queueSize := float64(CONFIGURATION.NUM_NIC_VM) * float64(NUM_VMs[i]) * CONFIGURATION.BUFF_SIZE
		virtual := float64(CONFIGURATION.NUM_NIC_VM) * float64(NUM_VMs[i]) * 80
		vmCapacity := float64(NUM_VMs[i]) * CONFIGURATION.VM_COMPUTE_CAP * float64(CONFIGURATION.NUM_NIC_VM)
		//dequeuePkts := int(math.Ceil(vmCapacity / (PKT_LEN * 1000)))
		dequeueBits := int(math.Ceil(vmCapacity / (CONFIGURATION.PROCESSING_DELAY)))
		INGRESS_CAP = append(INGRESS_CAP, new(VM))
		INGRESS_CAP[i].cap = vmCapacity
		INGRESS_CAP[i].vmQueue = queueSize
		//INGRESS_CAP[i].numOfDequeuePkts = dequeuePkts
		INGRESS_CAP[i].numOfDequeueBits = dequeueBits
		INGRESS_CAP[i].availableBuffSpace = queueSize

		_DEBUG.Printf("Function: initializeAdaptive - Initial capacity of ingress %d = %f", i, vmCapacity)
		//_DEBUG.Printf("Function: initializeAdaptive - Packets to dequeue at ingress %d = %d", i, dequeuePkts)
		_DEBUG.Printf("Function: initializeAdaptive - Bits to dequeue at ingress %d = %d", i, dequeueBits)
		_DEBUG.Printf("Function: initializeAdaptive - Initial queue size at ingress %d = %f", i, queueSize)
		_INFO.Printf("BufferSize %f ingress %d", virtual, i)

		// # CONFIGURATION.INGRESS_CAP[i] = math.floor(total_num_vms/CONFIGURATION.INGRESS_LOC)
	}
}

func diagnose(pkt packet) {
	diagnoseTraffic(pkt)
}

func mitigate() {
	if CONFIGURATION.DEFENSE_TYPE == "static" {
		staticMitigation()
	} else if CONFIGURATION.DEFENSE_TYPE == "dynamic" {
		dynamicMitigation()
	} else if CONFIGURATION.DEFENSE_TYPE == "adaptive" {
		adaptiveMitigation()
	}
}
