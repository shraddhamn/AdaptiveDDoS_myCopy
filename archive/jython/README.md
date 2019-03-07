# Adaptive - DDOS 

A Distributed Denial of Service (DDoS) defense system simulation which is capable of handling smart attackers. The DDoS defense system is adaptive and can allocate resources to ingress locations based on it's ability to predict the incoming attack traffic at each ingress by looking at the past traffic at each ingress location.


## Simulator Design 

The DDoS defense system has been designed to be configured in three states: Static, Dynamic and Adaptive. The static defense system is the basic defense system where there is there is no change in the number of VM's alloted to all the ingress locations. The dynamic defense system will allocate the the VM's for the ingress location without any prediction. The Adaptive defense system will decide the number of VM's to be allocated for each ingress location per epoch based on the past traffic at each ingress location. The assumptions made are, each virtual machine has one NIC and each NIC has one queue. We prove the efficiency of our defense model by computing the number of legitimate packets dropped and the proper utilization of allocated VM's to each ingress location. 


## This Repository Contains..

1. config.py -  The conditions necessary to set up the DDoS defense simulator.
2. packet.py -  The packet class containing information about the length, destination and protocol of the packet, the ingress              								 location it is sent to and also information on whether it is an attack packet or not. 
3. network.py -  Functionalities to define what happens when the packet arrives in the network.
4. buffer.py - Handles the enqueing and dequeing and dropping of packets. There is one queue per NIC of the VM. 
5. globals.py -  Declaration and Initialiazation of the global variables. 
6. mitigation_strategy.py - The three types of defense models : static, dynamic and adaptive have been defined here.
7. defense.py - the detection and the mitigation strategy have been defined here. 
8. detection_diagnosis.py - After detection whether the attack packet is a UDP or a TCP packet the functions for diagnosis of 														the type of attack(UDP flood, SYN flood, DNS amplication) are defined in this file.


## Requirement

Python 3.7.0

