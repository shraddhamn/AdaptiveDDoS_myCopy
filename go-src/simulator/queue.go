package main

func enqueue(pkt packet, i int) {
	pktQueue[i].Add(pkt)
}

func dequeue(i int) packet {
	return pktQueue[i].Next().(packet)
}