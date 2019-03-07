
import globals
import random, math
import logging
# this file contains the three various mitigation techniques
# that can be used by the defense layer, all traffic from the 
# diagnosis layer comes into the mitigation layer which 
# mitigates and creates flow rules etc.

# def load

# def static_mitigation():

	# add rules to controller to drop the attack traffic
	# tags traffic as attack or legitimate


# 1


	# for i in xrange(0, globals.INGRESS_LOC):
	# 	globals.NUM_VMs.append(math.floor(total_num_vms*1.0//globals.INGRESS_LOC))
	# 	queueSize = globals.NUM_NIC_VM*globals.NUM_VMs[i]*globals.BUFF_SIZE
	# 	dequeuePkts = globals.NUM_NIC_VM*globals.NUM_VMs[i]
	# 	vmCapacity = globals.NUM_VMs[i]*globals.VM_COMPUTE_CAP
	# 	globals.INGRESS_CAP.append(VM.VM(vmCapacity,queueSize,dequeuePkts,vmCapacity))


def changeCapacity(i,newCap):
	globals.NUM_VMs[i] = math.floor(newCap*1.0/globals.VM_COMPUTE_CAP)
	globals.LOCK_INGRESS_CAP[i].acquire()
	oldCap = globals.INGRESS_CAP[i].cap
	globals.INGRESS_CAP[i].cap = newCap
	globals.INGRESS_CAP[i].vmQueue = globals.NUM_NIC_VM*globals.NUM_VMs[i]*globals.BUFF_SIZE
	globals.INGRESS_CAP[i].numOfDequeuePkts = globals.NUM_NIC_VM*globals.NUM_VMs[i]
	globals.INGRESS_CAP[i].availableBuffSpace = newCap - (oldCap - globals.INGRESS_CAP[i].availableBuffSpace)
	# globals.DEBUG_LOGGER.debug(f"Function: change Capacity, capacity at ingress {i} = {newCap}")
	globals.LOCK_INGRESS_CAP[i].release()



def dynamicMitigation():

	for i in range(0,globals.INGRESS_LOC):
		
		globals.LOCK_CURR_TRAFFIC_STATS[i].acquire()
		if(globals.CURR_TRAFFIC_STATS[i]["total"] * globals.PKT_LEN > globals.PEAK_TRAFFIC[i]):
			globals.PEAK_TRAFFIC[i] = globals.CURR_TRAFFIC_STATUS[i]["total"] * globals.PKT_LEN


		if(globals.CURR_TRAFFIC_STATS[i]["total"]* globals.PKT_LEN < globals.MIN_TRAFFIC[i]):
			globals.MIN_TRAFFIC[i] = globals.CURR_TRAFFIC_STATS[i]["total"] * globals.PKT_LEN

		globals.LOCK_CURR_TRAFFIC_STATS[i].release()
		# globals.INGRESS_CAP[i] = random.uniform(globals.MIN_TRAFFIC[i],globals.PEAK_TRAFFIC[i])
		changeCapacity(i,random.uniform(globals.MIN_TRAFFIC[i],globals.PEAK_TRAFFIC[i]))
		

def adaptiveMitigation():

	for i in range(0,globals.INGRESS_LOC):
		
		globals.LOCK_CURR_TRAFFIC_STATS[i].acquire()
		if(globals.CURR_TRAFFIC_STATS[i]["total"] > globals.PEAK_TRAFFIC[i]):
			globals.PEAK_TRAFFIC[i] = globals.CURR_TRAFFIC_STATUS[i]["total"]


		if(globals.CURR_TRAFFIC_STATS[i]["total"] < globals.MIN_TRAFFIC[i]):
			globals.MIN_TRAFFIC[i] = globals.CURR_TRAFFIC_STATS[i]["total"]

		globals.LOCK_CURR_TRAFFIC_STATS[i].release()
		changeCapacity(i,random.uniform(globals.MIN_TRAFFIC[i],globals.PEAK_TRAFFIC[i]))

def staticMitigation():

	for i in range(0,globals.INGRESS_LOC):
		
		globals.LOCK_CURR_TRAFFIC_STATS[i].acquire()
		if(globals.CURR_TRAFFIC_STATS[i]["total"] > globals.PEAK_TRAFFIC[i]):
			globals.PEAK_TRAFFIC[i] = globals.CURR_TRAFFIC_STATUS[i]["total"]


		if(globals.CURR_TRAFFIC_STATS[i]["total"] < globals.MIN_TRAFFIC[i]):
			globals.MIN_TRAFFIC[i] = globals.CURR_TRAFFIC_STATS[i]["total"]

		globals.LOCK_CURR_TRAFFIC_STATS[i].release()
		changeCapacity(i,random.uniform(globals.MIN_TRAFFIC[i],globals.PEAK_TRAFFIC[i]))
		

	# set flow rules for new capacity in the controller
	# create packet rules for attack packers which is essentially the static mitigation


# def adaptive_mitigation(traffic_stats):


	



