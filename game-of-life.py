import os
import time
import random

# Create a 2D array to represent the game board
board = [[0 for x in range(80)] for y in range(40)]

# Initialize the game board with random values
for i in range(40):
    for j in range(80):
        board[i][j] = int(random.random() * 2)

# Main loop
while True:
    # Clear the screen
    os.system('clear')

    # Print the board
    for i in range(40):
        for j in range(80):
            if board[i][j] == 0:
                print('\033[40m \033[0m', end='')  # Black background
            else:
                print('\033[44m \033[0m', end='')  # Blue character
        print()

    # Calculate the next generation
    new_board = [[0 for x in range(80)] for y in range(40)]
    for i in range(40):
        for j in range(80):
            # Count the number of neighbors
            neighbors = 0
            for dx in range(-1, 2):
                for dy in range(-1, 2):
                    if dx == 0 and dy == 0:
                        continue
                    ni, nj = i + dx, j + dy
                    if ni >= 0 and ni < 40 and nj >= 0 and nj < 80 and board[ni][nj] == 1:
                        neighbors += 1

            # Apply the rules of the game
            if board[i][j] == 0 and neighbors == 3:
                new_board[i][j] = 1
            elif board[i][j] == 1 and (neighbors < 2 or neighbors > 3):
                new_board[i][j] = 0
            else:
                new_board[i][j] = board[i][j]

    # Update the board
    board = new_board

    # Wait for a bit
    time.sleep(0.1)

