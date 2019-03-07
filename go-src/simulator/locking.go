package main

import "sync"

var (
	LOCK_CURR_TRAFFIC_STATS []sync.Mutex
	// LOCK_RECEIVE_COUNTER []sync.Mutex
	LOCK_legitimateDropCounter []sync.Mutex
	LOCK_INGRESS_CAP []sync.Mutex
	// LOCK_NUM_VM []sync.Mutex
	// LOCK_RECEIVE_COUNTER []sync.Mutex
)

func initializeLocks() {

	LOCK_CURR_TRAFFIC_STATS = make([]sync.Mutex, CONFIGURATION.INGRESS_LOC)
	LOCK_INGRESS_CAP = make([]sync.Mutex, CONFIGURATION.INGRESS_LOC)
	// LOCK_NUM_VM = make([]sync.Mutex, CONFIGURATION.INGRESS_LOC)
	// LOCK_RECEIVE_COUNTER = make([]sync.Mutex, CONFIGURATION.INGRESS_LOC)
	LOCK_legitimateDropCounter = make([]sync.Mutex, CONFIGURATION.INGRESS_LOC)
	// for i := 0 ; i < CONFIGURATION.INGRESS_LOC ; i++ {
		


	// 	LOCK_CURR_TRAFFIC_STATS = append(LOCK_CURR_TRAFFIC_STATS,New(sync.Mutex))
	

	// 	// LOCK_legitimateDropCounter append(threading.Lock())
	// 	// globals.LOCK_RECEIVE_COUNTER.append(threading.Lock())

	// 	// # globals.LOCK_INGRESS_AVAILABLE_CAP.append(lock.Lock())
	// 	LOCK_INGRESS_CAP = append(LOCK_INGRESS_CAP, New(sync.Mutex))


	// 	LOCK_NUM_VM = append(LOCK_NUM_VM,New(sync.Mutex))
	// }

}



