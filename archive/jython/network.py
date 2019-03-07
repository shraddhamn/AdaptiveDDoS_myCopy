# Maybe add the propogation delay later 

import globals

import logging
from defense import *
import buffer

def sendtoNetwork(pkt):
	# pause till transmission delay
	receiveonNetwork(pkt)


def receiveonNetwork(pkt):
	deliverPacket(pkt)
	#detection part
	#call receive pkt


def deliverPacket(pkt):
	
	buffer.enqueuePacket(pkt)


