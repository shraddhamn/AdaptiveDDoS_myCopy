# This file contains the toplogy of the ISP
# including the number of ingress locations,
# the total capacity of the ISP

import globals
import math
import logging
import queue

prevDropCount = [0] * globals.INGRESS_LOC
prevReceiveCount = [0] * globals.INGRESS_LOC

def initializeISP():

	globals.DEBUG_LOGGER.debug("Function: initializeISP- Initializing TRAFFIC_STATS variables")

	for i in range(0,globals.INGRESS_LOC):
		# globals.CAPACITY.append(math.floor(globals.TOTAL_CAPACITY_ISP/globals.INGRESS_LOC))
		globals.CURR_TRAFFIC_STATS.append({})
		globals.CURR_TRAFFIC_STATS[i]["total"] = 0
		globals.CURR_TRAFFIC_STATS[i]["udp_flood"] = 0
		globals.CURR_TRAFFIC_STATS[i]["tcp_syn"] = 0

		globals.PREV_TRAFFIC_STATS.append({})
		globals.PREV_TRAFFIC_STATS[i]["total"] = 0
		globals.PREV_TRAFFIC_STATS[i]["udp_flood"] = 0
		globals.PREV_TRAFFIC_STATS[i]["tcp_syn"] = 0

		globals.PEAK_TRAFFIC.append(0)
		globals.MIN_TRAFFIC.append(float('inf'))
		globals.RECEIVE_COUNTER.append(0)
		globals.legitimateDropCounter.append(0)
		globals.processCounter.append(0)


		globals.BUFFER.append(queue.Queue())

		# globals.BUFF_SIZE.append(buff)



# def latencyIncurred():

def countDroppedPackets():

	global prevDropCount
	for i in range(0,globals.INGRESS_LOC):
		
		globals.LOCK_legitimateDropCounter[i].acquire()
		dropped = globals.legitimateDropCounter[i] - prevDropCount[i]
		prevDropCount[i] = globals.legitimateDropCounter[i]
		
		
		globals.DEBUG_LOGGER.debug(f"Function: countDroppedPackets, dropped Packets at ingress {i} are {dropped}")
		globals.STATS_LOGGER.info(f"Dropped Packets at ingress {i} = {dropped}")
		globals.LOCK_legitimateDropCounter[i].release()

	# loggings.error('This should go to both console and file')
 

def wastedResources():

	for i in range(0,globals.INGRESS_LOC):
		
		# globals.LOCK_RECEIVE_COUNTER[i].acquire()
		receivedPktsPerWIndow =  globals.RECEIVE_COUNTER[i] - prevReceiveCount[i]
		prevReceiveCount[i] = globals.RECEIVE_COUNTER[i] 
	

		globals.LOCK_INGRESS_CAP[i].acquire()
		wastedCap = globals.INGRESS_CAP[i].cap - (receivedPktsPerWIndow*globals.PKT_LEN)
		globals.LOCK_INGRESS_CAP[i].release()

		

		globals.DEBUG_LOGGER.debug(f"Function: wastedResources, wasted resources at ingress {i} are {wastedCap} Mbps")
		globals.STATS_LOGGER.info(f"Wasted resources at ingress {i} = {wastedCap} Mbps")

		# globals.LOCK_RECEIVE_COUNTER[i].release()

		# print wastedCap


