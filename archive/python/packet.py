# Packet definition
import globals
import logging
class Packet(object):
# Packet class contains the following:
#	1.length of the packet 
#	2.packet protocol - TCP/ UDP
#	3.packet destination - which ingress location it does to 

	def __init__(self, packet_len,protocol,ingress_loc, attack_flag):
		self.packet_len = packet_len
		self.protocol = protocol
		self.ingress = ingress_loc
		self.attack_flag = attack_flag
		self.dest = "dummy"
		self.detection = "benign"
		
