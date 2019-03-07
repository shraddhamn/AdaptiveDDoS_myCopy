


import globals
import detection_diagnosis
import mitigation_strategy
import VM
import math
import logging

def initializeAdaptive():
	globals.DEBUG_LOGGER.debug("Function: initializeAdaptive - Initialize inital capapcity of ingress locations")
	total_num_vms = math.floor(globals.ISP_CAP*1.0/globals.VM_COMPUTE_CAP)
	for i in range(0, globals.INGRESS_LOC):
		globals.NUM_VMs.append(math.floor(total_num_vms*1.0//globals.INGRESS_LOC))
		queueSize = globals.NUM_NIC_VM*globals.NUM_VMs[i]*globals.BUFF_SIZE
		dequeuePkts = globals.NUM_NIC_VM*globals.NUM_VMs[i]
		vmCapacity = globals.NUM_VMs[i]*globals.VM_COMPUTE_CAP
		globals.INGRESS_CAP.append(VM.VM(vmCapacity,queueSize,dequeuePkts,vmCapacity))
		globals.DEBUG_LOGGER.debug(f"Function: initializeAdaptive - Initial capacity of ingress {i} is {vmCapacity}")

		# globals.INGRESS_CAP[i] = math.floor(total_num_vms/globals.INGRESS_LOC)




def initialize():
	if(globals.DEFENSE_TYPE == "adaptive"):
		initializeAdaptive()


def diagnose(pkt):
	detection_diagnosis.diagnoseTraffic(pkt)



def mitigate():
	if(globals.DEFENSE_TYPE == "static"):
		mitigation_strategy.staticMitigation()
	elif(globals.DEFENSE_TYPE == "dynamic"):
		mitigation_strategy.dynamicMitigation()
	elif(globals.DEFENSE_TYPE == "adaptive"):
		mitigation_strategy.adaptiveMitigation()




