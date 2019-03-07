
from threading import Thread,Timer,Event, Lock
import logging
import globals
import log
import sys
import legitimate_traffic
import attack_traffic
import isp
import mylock
import buffer
import json
import defense
from io import open
import logging
import defense

u = str
class RepeatingThread(Thread):
    def __init__(self, event, t, name, m):
        Thread.__init__(self)
        self.stopped = event
        self.waitiingTime = t
        self.threadName = name
        self.method = m

    def run(self):
        globals.DEBUG_LOGGER.debug(f"Hello from thread {self.threadName}")
        globals.DEBUG_LOGGER.debug(f"Thread repeat time = {self.waitiingTime}")
        globals.DEBUG_LOGGER.debug(f"Thread Function = {self.method}")
        while not self.stopped.wait(self.waitiingTime):
            self.method()
            # print self.threadName
            # call a function


def startNewWindow():

	# defense.mitigate()
	globals.STATS_LOGGER.info(f"STATS FOR WINDOW {globals.WINDOW_COUNTER} - START")

	for i in range(0,globals.INGRESS_LOC):
		
		globals.LOCK_CURR_TRAFFIC_STATS[i].acquire()
		globals.PREV_TRAFFIC_STATS[i]["total"] = globals.CURR_TRAFFIC_STATS[i]["total"]
		globals.STATS_LOGGER.info(f"Total Traffic at Ingress {i} = {globals.CURR_TRAFFIC_STATS[i]['total']}")
		globals.CURR_TRAFFIC_STATS[i]["total"] = 0
		
		globals.PREV_TRAFFIC_STATS[i]["udp_flood"] = globals.CURR_TRAFFIC_STATS[i]["udp_flood"]
		globals.STATS_LOGGER.info(f"Total UDP Flood at Ingress {i} = {globals.CURR_TRAFFIC_STATS[i]['udp_flood']}")
		globals.CURR_TRAFFIC_STATS[i]["udp_flood"] = 0

		globals.PREV_TRAFFIC_STATS[i]["tcp_syn"] = globals.CURR_TRAFFIC_STATS[i]["tcp_syn"]
		globals.STATS_LOGGER.info(f"Total TCP Syn at Ingress {i} = {globals.CURR_TRAFFIC_STATS[i]['tcp_syn']}")
		globals.CURR_TRAFFIC_STATS[i]["tcp_syn"] = 0
		globals.LOCK_CURR_TRAFFIC_STATS[i].release()

	
	globals.WINDOW_COUNTER += 1
	isp.countDroppedPackets()
	isp.wastedResources()
	
	globals.STATS_LOGGER.info(f"STATS FOR WINDOW {globals.WINDOW_COUNTER} - END \n\n")


def readConfigureFile(file):
	data = ""
	with open(file, encoding="utf-8") as data_file:
		data = json.loads(data_file.read())

	# print data
	globals.ATTACK_TYPE = data[u("attackerType")]
	globals.DEFENSE_TYPE = data[u("defenseType")]
	globals.INGRESS_LOC = data[u("ingreeLoc")]
	globals.BUFF_SIZE = data[u("buffSize")] * 8 #conversion to Mbits (Mbytes)
	globals.VM_COMPUTE_CAP = data[u("VMCapacity")] * 1000 #conversion to Mbits (Gbps)
	globals.ISP_CAP = data[u("ISPCapacity")] * 1000 # conversion to Mbits (Gbps)
	globals.NUM_NIC_VM = data[u("numNICsVM")]
	globals.ATTACKER_CAP = data[u("attackerCapacity")] * 1000 #conversion to Mbits
	globals.LEG_TRAFFIC_MODEL = data[u("legitimateTraffic")]
	globals.EPOCH_TIME = data[u("epochTime")]	# seconds
	globals.PROCESSING_DELAY = data[u("processingDelay")]	# seconds

	globals.DEBUG_LOGGER.debug(f"Attack type = {globals.ATTACK_TYPE}")
	globals.DEBUG_LOGGER.debug(f"Defense type = {globals.DEFENSE_TYPE}")
	globals.DEBUG_LOGGER.debug(f"Number of Ingress Locations = {globals.INGRESS_LOC}")
	globals.DEBUG_LOGGER.debug(f"Buffer Size per queue = {globals.BUFF_SIZE} Mb")
	globals.DEBUG_LOGGER.debug(f"Computation capacity per VM  = {globals.VM_COMPUTE_CAP} Mbps")
	globals.DEBUG_LOGGER.debug(f"Total ISP capacity = {globals.ISP_CAP} Mbps")
	globals.DEBUG_LOGGER.debug(f"Number of NICs per VM = {globals.NUM_NIC_VM}")
	globals.DEBUG_LOGGER.debug(f"Total Attacker capacity = {globals.ATTACKER_CAP} Mbps")
	globals.DEBUG_LOGGER.debug(f"Legitimate traffic model = {globals.LEG_TRAFFIC_MODEL}")
	globals.DEBUG_LOGGER.debug(f"EPOCH duration = {globals.EPOCH_TIME} seconds")
	globals.DEBUG_LOGGER.debug(f"Processing delay = {globals.PROCESSING_DELAY} seconds")




def main():
	

	# read conf file to set globals
	log.loggingSetup()
	
	globals.DEBUG_LOGGER.debug('Hello from main')

	globals.DEBUG_LOGGER.debug('Reading configuration file')
	configurationFile = sys.argv[1]
	
	readConfigureFile(configurationFile)


	# logging conf
	
	globals.DEBUG_LOGGER.debug('Initialize traffic stats data structures in ISP')
	# initialize traffic stats data structures
	isp.initializeISP()

	# initialize capacity of ingress locations depending on the defense type
	globals.DEBUG_LOGGER.debug('Initialize capacity of ingress locations depending on the defense type')
	defense.initialize()


	mylock.initializeLocks()
	# start legitimate traffic thread
	globals.DEBUG_LOGGER.debug('Start legitimate traffic thread')
	Thread(target=legitimate_traffic.flowGen).start()

	# start attack traffic thread
	globals.DEBUG_LOGGER.debug('Start attack traffic thread')
	# threading.Thread(target=attackTrafficGen, args=[attackerType]).start()
	


	# start per epoch stat collection for next epoch prediction
	globals.DEBUG_LOGGER.debug('Start stats collection thread')
	stopStats = Event()
	statsThread = RepeatingThread(stopStats,globals.EPOCH_TIME,"stats per epoch collection thread",startNewWindow)
	statsThread.start()


	# start dequeuing pkts after delay equivalent to processing delay
	globals.DEBUG_LOGGER.debug("Start packet processing thread")
	stopProcess = Event()
	processingThread = RepeatingThread(stopProcess,globals.PROCESSING_DELAY,"packet processingThread",buffer.processPacket)
	processingThread.start()

#	 this will stop the timer
	# stopFlag.set()


if __name__ == "__main__":
    main()


