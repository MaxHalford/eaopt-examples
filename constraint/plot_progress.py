import json
import math
import os.path
import sys

from matplotlib import animation
import matplotlib.pyplot as plt
import numpy as np


PROGRESS_FILE_NAME = 'progress.json'
BOUNDS = [-10, 10]


def parabola(x, y):
    return x ** 2 + y ** 2


def main():

    # Check the progress file exists
    if not os.path.isfile(PROGRESS_FILE_NAME):
        print('{} does not exit, execute "go run main.go" first'.format(PROGRESS_FILE_NAME))
        sys.exit()

    # Load each line in the progress file
    steps = [json.loads(line) for line in open(PROGRESS_FILE_NAME)]

    # Extract the coordinates of each individual at each step
    coordinates = np.array([
        [
            indi['genome']
            for indi in step['pops'][0]['indis']
        ]
        for step in steps
    ])

    fig = plt.figure(figsize=(8, 8))
    ax = plt.axes(xlim=BOUNDS, ylim=BOUNDS)

    # Draw a heatmap of the function
    X = np.linspace(BOUNDS[0], BOUNDS[1], 500)
    Y = np.linspace(BOUNDS[0], BOUNDS[1], 500)
    Z = parabola(*np.meshgrid(X, Y))
    heatmap = ax.imshow(Z, extent=BOUNDS * 2)

    # Plot the constraint
    x = np.arange(-4, 4.1, 0.1)
    ax.fill_between(x, 4 - np.abs(x), np.abs(x) - 4, facecolor='gray', alpha=0.8)

    line, = ax.plot([], [], marker='o', color='orange', linestyle='None')
    best, = ax.plot([], [], marker='*', color='red', linestyle='None', markersize=30)

    def init():
        line.set_data([], [])
        best.set_data([], [])
        return line, best

    def animate(i):
        line.set_data(coordinates[i][1:, 0], coordinates[i][1:, 1])
        best.set_data([coordinates[i][0, 0]], [coordinates[i][0, 1]])
        return line, best

    anim = animation.FuncAnimation(fig, animate, init_func=init,
                               frames=len(coordinates), interval=100, blit=True)

    try:
        print('Saving to GIF')
        anim.save('progress.gif', writer='imagemagick', fps=10)
    except:
        print('Could not save progress to GIF')

    plt.show()


if __name__ == '__main__':
    main()
