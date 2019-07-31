#!/usr/bin/env python

import matplotlib.pyplot as plt
import numpy as np


p = (1013,1513,3038)
y_pos = np.arange(len(p))
#FABRIC_BASE = RED
#FABRIC-SOLO = BLUE
#FABRIC-TLS = GREEN
#FABRIC-PRIVATE = CYAN
fig, axs = plt.subplots(4, 3, constrained_layout=True)

fabricBase =axs[0][0].barh(y_pos+0.1, [2.74,2.23,1.86],height=0.2, color="red")


fabricSolo =axs[0][0].barh(y_pos-0.1, [2.69,2.39,1.78],height=0.2, color='blue')

fabricTLS =axs[0][0].barh(y_pos-0.3, [0.71,0.56,0.46],height=0.2, color='green')

fabricPrivate =axs[0][0].barh(y_pos+0.3, [24.88,34.26,20.41],height=0.2, color='cyan')

axs[0][0].set_title('5 clients 25tps')
axs[0][0].set_ylabel('Payload in bytes')
axs[0][0].set_xlabel('Average latency (seconds)')
axs[0][0].set_yticklabels(p)
axs[0][0].set_yticks(y_pos)
axs[0][0].invert_yaxis()  

#FABRIC_BASE = RED
#FABRIC-SOLO = BLUE
#FABRIC-TLS = GREEN
#FABRIC-PRIVATE = CYAN

axs[1][0].set_title('5 clients 50tps')
axs[1][0].barh(y_pos+0.1, [2.92,3.03,3.40],height=0.2, color="red")
axs[1][0].barh(y_pos-0.1, [2.32,2.21,1.62],height=0.2, color='blue')
axs[1][0].barh(y_pos-0.3, [0.75,6.72,4.04],height=0.2, color='green')
axs[1][0].barh(y_pos+0.3, [35.18,34.08,33.04],height=0.2, color='cyan')

axs[1][0].set_ylabel('Payload in bytes')
axs[1][0].set_xlabel('Average latency (seconds)')
axs[1][0].set_yticklabels(p)
axs[1][0].set_yticks(y_pos)
axs[1][0].invert_yaxis() 

#FABRIC_BASE = RED
#FABRIC-SOLO = BLUE
#FABRIC-TLS = GREEN
#FABRIC-PRIVATE = CYAN

axs[2][0].set_title('5 clients 100tps')
axs[2][0].barh(y_pos+0.1, [8.15,6.49,8.90],height=0.2, color="red")
axs[2][0].barh(y_pos-0.1, [7.65,5.61,7.60],height=0.2, color='blue')
axs[2][0].barh(y_pos-0.3, [11.12,9.73,9.69],height=0.2, color='green')
axs[2][0].barh(y_pos+0.3, [34.35,36.16,36.93],height=0.2, color='cyan')

axs[2][0].set_ylabel('Payload in bytes')
axs[2][0].set_xlabel('Average latency (seconds)')
axs[2][0].set_yticklabels(p)
axs[2][0].set_yticks(y_pos)
axs[2][0].invert_yaxis() 

#FABRIC_BASE = RED
#FABRIC-SOLO = BLUE
#FABRIC-TLS = GREEN
#FABRIC-PRIVATE = CYAN

axs[3][0].set_title('5 clients 1000tps')
axs[3][0].barh(y_pos+0.1, [13.33,12.95,12.39],height=0.2, color="red")
axs[3][0].barh(y_pos-0.1, [10.63,12.20,13.88],height=0.2, color='blue')
axs[3][0].barh(y_pos-0.3, [15.69,13.81,16.16],height=0.2, color='green')
axs[3][0].barh(y_pos+0.3, [40.47,36.80,37.93],height=0.2, color='cyan')

axs[3][0].set_ylabel('Payload in bytes')
axs[3][0].set_xlabel('Average latency (seconds)')
axs[3][0].set_yticklabels(p)
axs[3][0].set_yticks(y_pos)
axs[3][0].invert_yaxis() 

#FABRIC_BASE = RED
#FABRIC-SOLO = BLUE
#FABRIC-TLS = GREEN
#FABRIC-PRIVATE = CYAN

axs[0][1].set_title('10 clients 25tps')
axs[0][1].barh(y_pos+0.1, [2.69,2.26,1.81],height=0.2, color="red")
axs[0][1].barh(y_pos-0.1, [2.68,2.40,1.79],height=0.2, color='blue')
axs[0][1].barh(y_pos-0.3, [0.71,0.59,0.46],height=0.2, color='green')
axs[0][1].barh(y_pos+0.3, [22.34,23.74,19.55],height=0.2, color='cyan')

axs[0][1].set_ylabel('Payload in bytes')
axs[0][1].set_xlabel('Average latency (seconds)')
axs[0][1].set_yticklabels(p)
axs[0][1].set_yticks(y_pos)
axs[0][1].invert_yaxis() 

#FABRIC_BASE = RED
#FABRIC-SOLO = BLUE
#FABRIC-TLS = GREEN
#FABRIC-PRIVATE = CYAN

axs[1][1].set_title('10 clients 50tps')
axs[1][1].barh(y_pos+0.1, [2.43,2.26,2.45],height=0.2, color="red")
axs[1][1].barh(y_pos-0.1, [2.00,2.01,1.71],height=0.2, color='blue')
axs[1][1].barh(y_pos-0.3, [3.01,1.26,4.75],height=0.2, color='green')
axs[1][1].barh(y_pos+0.3, [36.12,34.06,33.50],height=0.2, color='cyan')

axs[1][1].set_ylabel('Payload in bytes')
axs[1][1].set_xlabel('Average latency (seconds)')
axs[1][1].set_yticklabels(p)
axs[1][1].set_yticks(y_pos)
axs[1][1].invert_yaxis() 

#FABRIC_BASE = RED
#FABRIC-SOLO = BLUE
#FABRIC-TLS = GREEN
#FABRIC-PRIVATE = CYAN

axs[2][1].set_title('10 clients 100tps')
axs[2][1].barh(y_pos+0.1, [7.12,8.51,8.73],height=0.2, color="red")
axs[2][1].barh(y_pos-0.1, [3.53,5.68,6.30],height=0.2, color='blue')
axs[2][1].barh(y_pos-0.3, [7.18,6.18,4.63],height=0.2, color='green')
axs[2][1].barh(y_pos+0.3, [34.04,35.95,34.87],height=0.2, color='cyan')

axs[2][1].set_ylabel('Payload in bytes')
axs[2][1].set_xlabel('Average latency (seconds)')
axs[2][1].set_yticklabels(p)
axs[2][1].set_yticks(y_pos)
axs[2][1].invert_yaxis() 
#FABRIC_BASE = RED
#FABRIC-SOLO = BLUE
#FABRIC-TLS = GREEN
#FABRIC-PRIVATE = CYAN

axs[3][1].set_title('10 clients 1000tps')
axs[3][1].barh(y_pos+0.1, [13.69,14.76,12.36],height=0.2, color="red")
axs[3][1].barh(y_pos-0.1, [9.42,13.11,12.80],height=0.2, color='blue')
axs[3][1].barh(y_pos-0.3, [13.04,13.24,11.06],height=0.2, color='green')
axs[3][1].barh(y_pos+0.3, [39.11,41.77,38.04],height=0.2, color='cyan')

axs[3][1].set_ylabel('Payload in bytes')
axs[3][1].set_xlabel('Average latency (seconds)')
axs[3][1].set_yticklabels(p)
axs[3][1].set_yticks(y_pos)
axs[3][1].invert_yaxis() 

#FABRIC_BASE = RED
#FABRIC-SOLO = BLUE
#FABRIC-TLS = GREEN
#FABRIC-PRIVATE = CYAN

axs[0][2].set_title('20 clients 25tps')
axs[0][2].barh(y_pos+0.1, [2.71,2.46,1.83],height=0.2, color="red")
axs[0][2].barh(y_pos-0.1, [2.60,2.34,1.77],height=0.2, color='blue')
axs[0][2].barh(y_pos-0.3, [0.68,0.59,0.47],height=0.2, color='green')
axs[0][2].barh(y_pos+0.3, [21.74,20.72,21.08],height=0.2, color='cyan')

axs[0][2].set_ylabel('Payload in bytes')
axs[0][2].set_xlabel('Average latency (seconds)')
axs[0][2].set_yticklabels(p)
axs[0][2].set_yticks(y_pos)
axs[0][2].invert_yaxis() 

#FABRIC_BASE = RED
#FABRIC-SOLO = BLUE
#FABRIC-TLS = GREEN
#FABRIC-PRIVATE = CYAN


axs[1][2].set_title('20 clients 50tps')
axs[1][2].barh(y_pos+0.1, [2.41,2.78,2.26],height=0.2, color="red")
axs[1][2].barh(y_pos-0.1, [2.53,1.86,1.75],height=0.2, color='blue')
axs[1][2].barh(y_pos-0.3, [3.10,3.96,6.63],height=0.2, color='green')
axs[1][2].barh(y_pos+0.3, [34.07,32.38,34.50],height=0.2, color='cyan')

axs[1][2].set_ylabel('Payload in bytes')
axs[1][2].set_xlabel('Average latency (seconds)')
axs[1][2].set_yticklabels(p)
axs[1][2].set_yticks(y_pos)
axs[1][2].invert_yaxis() 

#FABRIC_BASE = RED
#FABRIC-SOLO = BLUE
#FABRIC-TLS = GREEN
#FABRIC-PRIVATE = CYAN

axs[2][2].set_title('20 clients 100tps')
axs[2][2].barh(y_pos+0.1, [8.13,4.92,4.83],height=0.2, color="red")
axs[2][2].barh(y_pos-0.1, [5.68,5.21,9.45],height=0.2, color='blue')
axs[2][2].barh(y_pos-0.3, [6.92,7.13,9.47],height=0.2, color='green')
axs[2][2].barh(y_pos+0.3, [35.63,35.62,35.71],height=0.2, color='cyan')

axs[2][2].set_ylabel('Payload in bytes')
axs[2][2].set_xlabel('Average latency (seconds)')
axs[2][2].set_yticklabels(p)
axs[2][2].set_yticks(y_pos)
axs[2][2].invert_yaxis() 

#FABRIC_BASE = RED
#FABRIC-SOLO = BLUE
#FABRIC-TLS = GREEN
#FABRIC-PRIVATE = CYAN

axs[3][2].set_title('20 clients 1000tps')
axs[3][2].barh(y_pos+0.1, [13.70,12.25,11.95],height=0.2, color="red")
axs[3][2].barh(y_pos-0.1, [11.34,10.17,9.98],height=0.2, color='blue')
axs[3][2].barh(y_pos-0.3, [8.46,15.62,16.01],height=0.2, color='green')
axs[3][2].barh(y_pos+0.3, [39.30,41.42,40.33],height=0.2, color='cyan')

axs[3][2].set_ylabel('Payload in bytes')
axs[3][2].set_xlabel('Average latency (seconds)')
axs[3][2].set_yticklabels(p)
axs[3][2].set_yticks(y_pos)
axs[3][2].invert_yaxis() 


fig.legend((fabricBase,fabricPrivate,fabricSolo,fabricTLS),("Fabric-base","Fabric-private","Fabric-solo","Fabric-tls"),"lower center",ncol=4,bbox_to_anchor=(0.5, -0.01))
mng = plt.get_current_fig_manager()
mng.full_screen_toggle()
plt.savefig("./FigureTest.pdf",bbox_inches='tight')
plt.show()

