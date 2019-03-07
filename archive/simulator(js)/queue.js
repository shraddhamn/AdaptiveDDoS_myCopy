import time 
import packet 
import threading

lock = threading.Lock()


def receivepkt(pkt):
   if ((BUFFSIZE - pkt.packet_len) > 0):
      print_lock.acquire()
      receiveCounter+=1
      BUFFSIZE -= pkt.packet_len
      print_lock.release()
   else:
      droppkt()


def processedpkt(pkt,processingdelay):
   time.sleep(1/processingdelay)
   BUFFSIZE += pcket_len
   processCounter+=1

def droppkt(pkt):
   if (pkt.attack_flag == 0):
      legitimate_dropCounter+=1
   else:
      attack_dropCounter+=1

