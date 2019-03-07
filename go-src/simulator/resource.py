


import sys

total = 20000.0
filename = sys.argv[1]
f = open(filename,'r')
resources = 0
for line in f:
    if "Wasted_resources_Mbps" in line:
        line = line.strip().split(" ")
        res = float(line[5])
        if res > 0:
            resources += float(line[5])



avg = resources / 50
print ((avg/total)*100)




