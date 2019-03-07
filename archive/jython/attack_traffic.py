
# this file will be used to generate attacker strategies and send attack traffic
from globals import *
import logging
import random
import time
import globals 
from network import * 

def UDP_flood(pkt):
	globals.DEBUG_LOGGER.debug("Function: UDP_flood")
	for i in range (0, globals.INGRESS_LOC):
		fixedRate = 1000000
		numPkts = int(fixedRate / globals.PKT_LEN)
		for k in range (0, numPkts):
			network.sendtoNetwork(packet.Packet(globals.PKT_LEN,"udp",i,1))
			time.sleep(1)
	
	
		
		
	

	
