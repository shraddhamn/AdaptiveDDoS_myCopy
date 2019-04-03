package main

func enqueue(pkt packet, i int) {
	pktQueue[i].Add(pkt)
}

func dequeue(i int) packet {
	pkt := pktQueue[i].Next()
	if pkt == nil {
		pkt = NewPacket(0, "nil", -1, false)
	}
	return pkt.(packet)
}
