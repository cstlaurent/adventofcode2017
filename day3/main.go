package main

import "fmt"

const value1 = 361527

func main() {
	fmt.Printf("Day 3 moves: %d\n", calcMoves(value1))
	fmt.Printf("Day 3 neighbors: %d\n", calcNeighbors(value1))
}

const xDir = 0
const yDir = 1

func calcNeighbors(input int) int {
	off := 50
	t := [100][100]int{}
	t[50][50] = 1
	posX := 0
	posY := 0
	diam := 0
	currentVal := 1
	nextDir := []int{1, 0}
	for {
		posX += nextDir[xDir]
		posY += nextDir[yDir]

		// Top right corner
		if posX >= diam && posY >= diam {
			nextDir[xDir] = -1
			nextDir[yDir] = 0
		}

		// Top left corner
		if posX <= -diam && posY >= diam {
			nextDir[xDir] = 0
			nextDir[yDir] = -1
		}

		// Bottom left corner
		if posX <= -diam && posY <= -diam {
			nextDir[xDir] = 1
			nextDir[yDir] = 0
		}

		// Bottom right corner
		if posX > diam && posY <= -diam {
			nextDir[xDir] = 0
			nextDir[yDir] = 1
			diam++
		}
		currentVal = getNSum(t, posX, posY)
		fmt.Println("Val ", currentVal, " for pos X ", posX+off, ", Y ", posY+off)
		t[posX+off][posY+off] = currentVal
		if currentVal >= value1 {
			break
		}
	}
	return currentVal
}

func getNSum(t [100][100]int, posX int, posY int) int {
	oX := posX + 50
	oY := posY + 50
	return t[oX+1][oY] +
		t[oX+1][oY+1] +
		t[oX][oY+1] +
		t[oX-1][oY+1] +
		t[oX-1][oY] +
		t[oX-1][oY-1] +
		t[oX][oY-1] +
		t[oX+1][oY-1]
}

func calcMoves(input int) int {
	posX := 0
	posY := 0
	diam := 0
	currentVal := 1
	nextDir := []int{1, 0}
	for {
		posX += nextDir[xDir]
		posY += nextDir[yDir]

		// Top right corner
		if posX >= diam && posY >= diam {
			nextDir[xDir] = -1
			nextDir[yDir] = 0
		}

		// Top left corner
		if posX <= -diam && posY >= diam {
			nextDir[xDir] = 0
			nextDir[yDir] = -1
		}

		// Bottom left corner
		if posX <= -diam && posY <= -diam {
			nextDir[xDir] = 1
			nextDir[yDir] = 0
		}

		// Bottom right corner
		if posX > diam && posY <= -diam {
			nextDir[xDir] = 0
			nextDir[yDir] = 1
			diam++
		}
		currentVal++
		if currentVal >= value1 {
			break
		}
	}
	return posX + posY
}
