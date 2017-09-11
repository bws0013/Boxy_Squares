What is the minimum number of squares that must be filled in on a sudoku board before a 'dumb' program for solving it can solve it?

This program is part 1 of this problem. This is the required dumb program.

Each square of the board will run on its own thread and will constantly check if it is certain of its own value. Once the board fills and/or the max number of squares have been filled with their value the simulation ends.

At the start of execution there should be 81 squares each constantly checking themselves as well as the board itself checking to see if it is full yet. When one knows its number it should assign it to itself so that the other squares can see it and then quit.

An unassigned square is marked with a 0, an assigned square is marked with a number from 1 to 9.

Boards are located in storage. They consist of a 9x9 .csv file.

### In progress
