package main

import "fmt"

const value1 = 361527

func main() {
	fmt.Printf("Day 3 moves: %d\n", calcMoves(value1))
}

const xDir = 0
const yDir = 1

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
