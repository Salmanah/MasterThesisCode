import matplotlib.pyplot as plt

# x-coordinates of left sides of bars
left = [1, 2, 3, 4, 5]

# heights of bars
height = [100, 250, 500, 750, 1000]

# labels for bars
tick_label = [5, 10, 20]

# plotting a bar chart
plt.bar(left, height, tick_label, color = ['red', 'green'])

# naming the x-axis
plt.xlabel('Clients')
# naming the y-axis
plt.ylabel('Send rate')
# plot title
plt.title('My bar chart!')

# function to show the plot
plt.show()
