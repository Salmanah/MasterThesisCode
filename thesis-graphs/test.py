#!/usr/bin/env python

import matplotlib.pyplot as plt
import numpy as np


p = (100,250,500)
y_pos = np.arange(len(p))
#FABRIC_BASE = RED
#FABRIC-SOLO = BLUE
#FABRIC-TLS = GREEN
#FABRIC-PRIVATE = CYAN
fig, axs = plt.subplots(4, 3, constrained_layout=True)

fabricBase =axs[0][0].barh(y_pos+0.1, [2.70,2.93,2.73],height=0.2, color="red")
fabricSolo =axs[0][0].barh(y_pos-0.1, [2.71,2.82,2.62],height=0.2, color='blue')
fabricTLS =axs[0][0].barh(y_pos-0.3, [2.86,2.91,2.69],height=0.2, color='green')
fabricPrivate =axs[0][0].barh(y_pos+0.3, [13.10,29.53,34.68],height=0.2, color='cyan')

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
axs[1][0].barh(y_pos-0.1, [2.01,1.98,1.84],height=0.2, color='blue')
axs[1][0].barh(y_pos-0.3, [1.89,1.84,1.93],height=0.2, color='green')
axs[1][0].barh(y_pos+0.3, [24.57,21.91,32.00],height=0.2, color='cyan')

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
axs[2][0].barh(y_pos-0.1, [4.48,4.83,2.70],height=0.2, color='blue')
axs[2][0].barh(y_pos-0.3, [2.93,3.46,2.42],height=0.2, color='green')
axs[2][0].barh(y_pos+0.3, [27.53,26.93,26.24],height=0.2, color='cyan')

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
axs[3][0].barh(y_pos-0.1, [7.71,8.87,8.11],height=0.2, color='blue')
axs[3][0].barh(y_pos-0.3, [9.52,8.54,8.97],height=0.2, color='green')
axs[3][0].barh(y_pos+0.3, [34.62,29.53,34.68],height=0.2, color='cyan')

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
axs[0][1].barh(y_pos-0.1, [2.71,2.87,2.62],height=0.2, color='blue')
axs[0][1].barh(y_pos-0.3, [2.73,2.84,2.66],height=0.2, color='green')
axs[0][1].barh(y_pos+0.3, [21.10,8.21,9.89],height=0.2, color='cyan')

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
axs[1][1].barh(y_pos-0.1, [1.84,1.94,1.75],height=0.2, color='blue')
axs[1][1].barh(y_pos-0.3, [1.98,1.98,1.82],height=0.2, color='green')
axs[1][1].barh(y_pos+0.3, [22.56,20.62,24.41],height=0.2, color='cyan')

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
axs[2][1].barh(y_pos-0.1, [2.18,1.96,1.41],height=0.2, color='blue')
axs[2][1].barh(y_pos-0.3, [2.38,2.26,3.42],height=0.2, color='green')
axs[2][1].barh(y_pos+0.3, [31.12,26.57,25.81],height=0.2, color='cyan')

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
axs[3][1].barh(y_pos-0.1, [7.05,6.31,4.63],height=0.2, color='blue')
axs[3][1].barh(y_pos-0.3, [8.97,7.92,8.08],height=0.2, color='green')
axs[3][1].barh(y_pos+0.3, [32.78,29.67,28.92],height=0.2, color='cyan')

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
axs[0][2].barh(y_pos-0.1, [2.44,2.79,2.66],height=0.2, color='blue')
axs[0][2].barh(y_pos-0.3, [2.72,2.88,2.63],height=0.2, color='green')
axs[0][2].barh(y_pos+0.3, [16.87,15.85,12.49],height=0.2, color='cyan')

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
axs[1][2].barh(y_pos-0.1, [1.72,1.92,1.83],height=0.2, color='blue')
axs[1][2].barh(y_pos-0.3, [1.98,2.42,1.82],height=0.2, color='green')
axs[1][2].barh(y_pos+0.3, [25.26,25.68,24.38],height=0.2, color='cyan')

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
axs[2][2].barh(y_pos-0.1, [3.18,1.65,2.40],height=0.2, color='blue')
axs[2][2].barh(y_pos-0.3, [1.70,2.71,1.51],height=0.2, color='green')
axs[2][2].barh(y_pos+0.3, [26.74,36.69,25.10],height=0.2, color='cyan')

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
axs[3][2].barh(y_pos-0.1, [7.04,6.71,6.46],height=0.2, color='blue')
axs[3][2].barh(y_pos-0.3, [6.74,9.38,6.72],height=0.2, color='green')
axs[3][2].barh(y_pos+0.3, [28.87,43.05,31.52],height=0.2, color='cyan')

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

