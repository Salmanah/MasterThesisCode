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

fabricBase =axs[0][0].barh(y_pos+0.1, [2.70,2.93,2.73],height=0.2, color="red")
fabricSolo =axs[0][0].barh(y_pos-0.1, [2.69,2.39,1.78],height=0.2, color='blue')
fabricTLS =axs[0][0].barh(y_pos-0.3, [2.86,2.91,2.69],height=0.2, color='green')
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
axs[1][0].barh(y_pos+0.1, [1.97,2.1,2.11],height=0.2, color="red")
axs[1][0].barh(y_pos-0.1, [2.32,2.21,1.62],height=0.2, color='blue')
axs[1][0].barh(y_pos-0.3, [1.89,1.84,1.93],height=0.2, color='green')
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
axs[2][0].barh(y_pos+0.1, [3.62,3.19,4.38],height=0.2, color="red")
axs[2][0].barh(y_pos-0.1, [7.65,5.61,7.60],height=0.2, color='blue')
axs[2][0].barh(y_pos-0.3, [2.93,3.46,2.42],height=0.2, color='green')
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
axs[3][0].barh(y_pos+0.1, [8.99,8.21,10.06],height=0.2, color="red")
axs[3][0].barh(y_pos-0.1, [10.63,12.20,13.88],height=0.2, color='blue')
axs[3][0].barh(y_pos-0.3, [9.52,8.54,8.97],height=0.2, color='green')
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
axs[0][1].barh(y_pos+0.1, [2.75,2.79,2.76],height=0.2, color="red")
axs[0][1].barh(y_pos-0.1, [2.68,2.40,1.79],height=0.2, color='blue')
axs[0][1].barh(y_pos-0.3, [2.73,2.84,2.66],height=0.2, color='green')
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
axs[1][1].barh(y_pos+0.1, [1.93,1.99,2.16],height=0.2, color="red")
axs[1][1].barh(y_pos-0.1, [2.00,2.01,1.71],height=0.2, color='blue')
axs[1][1].barh(y_pos-0.3, [1.98,1.98,1.82],height=0.2, color='green')
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
axs[2][1].barh(y_pos+0.1, [1.83,2.49,3.43],height=0.2, color="red")
axs[2][1].barh(y_pos-0.1, [3.53,5.68,6.30],height=0.2, color='blue')
axs[2][1].barh(y_pos-0.3, [2.38,2.26,3.42],height=0.2, color='green')
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
axs[3][1].barh(y_pos+0.1, [7.66,9.70,7.37],height=0.2, color="red")
axs[3][1].barh(y_pos-0.1, [9.42,13.11,12.80],height=0.2, color='blue')
axs[3][1].barh(y_pos-0.3, [8.97,7.92,8.08],height=0.2, color='green')
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
axs[0][2].barh(y_pos+0.1, [2.70,2.98,2.75],height=0.2, color="red")
axs[0][2].barh(y_pos-0.1, [2.60,2.34,1.77],height=0.2, color='blue')
axs[0][2].barh(y_pos-0.3, [2.72,2.88,2.63],height=0.2, color='green')
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
axs[1][2].barh(y_pos+0.1, [1.93,2.17,2.05],height=0.2, color="red")
axs[1][2].barh(y_pos-0.1, [2.53,1.86,1.75],height=0.2, color='blue')
axs[1][2].barh(y_pos-0.3, [1.98,2.42,1.82],height=0.2, color='green')
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
axs[2][2].barh(y_pos+0.1, [3.39,3.37,1.83],height=0.2, color="red")
axs[2][2].barh(y_pos-0.1, [5.68,5.21,9.45],height=0.2, color='blue')
axs[2][2].barh(y_pos-0.3, [1.70,2.71,1.51],height=0.2, color='green')
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
axs[3][2].barh(y_pos+0.1, [8.81,6.48,8.70],height=0.2, color="red")
axs[3][2].barh(y_pos-0.1, [11.34,10.17,9.98],height=0.2, color='blue')
axs[3][2].barh(y_pos-0.3, [6.74,9.38,6.72],height=0.2, color='green')
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

