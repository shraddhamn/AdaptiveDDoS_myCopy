



# PROPAGATION_DELAY = 10
# TRANSMISSION_DELAY = 0
# PROCESSING_DELAY = 2
import logging


ATTACK_TYPE = "randIngress"
DEFENSE_TYPE = "adaptive"
INGRESS_LOC = 10
BUFF_SIZE = 100
VM_COMPUTE_CAP = 80
ISP_CAP = 90
NUM_PORTS_VM = 10
ATTACKER_CAP = 60
LEG_TRAFFIC_MODEL = "simple"
EPOCH_TIME = 60
PROCESSING_DELAY = 0.1

WINDOW_COUNTER = 0


TCP_SYN_DETECT_ACCURACY = 0.9
UDP_DETECT_ACCURACY = 0.9
#BUFFER_SIZE
# #-----QUEUING_DELAY = PROCESSING_DELAY*NUM_PKTS_IN_QUEUE

NUM_VMs = []
INGRESS_CAP = []

PREV_TRAFFIC_STATS = []
CURR_TRAFFIC_STATS = []
PEAK_TRAFFIC = []
MIN_TRAFFIC = []

PKT_LEN = 5
# UDP_DETECTION = 0.9

RECEIVE_COUNTER = [0] * INGRESS_LOC
legitimateDropCounter = [0] * INGRESS_LOC
processCounter = [0] * INGRESS_LOC


STATS_LOGGER = logging.getLogger('statsLogger')
DEBUG_LOGGER = logging.getLogger('debugLogger')




