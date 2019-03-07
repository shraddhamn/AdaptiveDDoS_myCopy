


import sys

total = 2000.0
filename = sys.argv[1]
f = open(filename,'r')
resources = 0
for line in f:
    if "Wasted_resources_Mbps" in line:
        line = line.strip().split(" ")
        res = float(line[5])
        if res < 0:
            a = math.abs(res)
            if a > 8:
                diff = a - 8
                pkts = 




avg = resources / 50
print ((avg/total)*100)




