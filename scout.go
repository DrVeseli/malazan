package main

import (
	"math/rand"
	"time"
)

type Direction struct {
	x, y int
}

var directions = []Direction{
	{-1, 0}, // left
	{1, 0},  // right
	{0, -1}, // up
	{0, 1},  // down
}

func findPlayerPosition() (int, int) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if worldMap[i*100+j] == 1 {
				return i, j
			}
		}
	}
	return -1, -1
}

func isValidMove(x, y int) bool {
	if x < 0 || x >= 100 || y < 0 || y >= 100 {
		return false
	}
	value := worldMap[x*100+y]
	return value == 0 || value == 3
}

func spawnScout() {
	playerX, playerY := findPlayerPosition()
	if playerX == -1 {
		return
	}

	// Try to find a valid spawn point around the player
	var scoutX, scoutY int
	validSpawn := false

	for _, dir := range directions {
		newX, newY := playerX+dir.x, playerY+dir.y
		if isValidMove(newX, newY) {
			scoutX, scoutY = newX, newY
			validSpawn = true
			break
		}
	}

	if !validSpawn {
		return
	}

	// Set scout position
	worldMap[scoutX*100+scoutY] = 2
	SaveWorldMap(worldMap)

	// Start scout movement
	go moveScout(scoutX, scoutY)
}

func moveScout(x, y int) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// Clear current position
		worldMap[x*100+y] = 3

		// Find valid moves
		var validMoves []Direction
		for _, dir := range directions {
			newX, newY := x+dir.x, y+dir.y
			if isValidMove(newX, newY) {
				validMoves = append(validMoves, dir)
			}
		}

		if len(validMoves) == 0 {
			SaveWorldMap(worldMap)
			return
		}

		// Choose random valid direction
		dir := validMoves[rand.Intn(len(validMoves))]
		x, y = x+dir.x, y+dir.y

		// Set new position
		worldMap[x*100+y] = 2

		SaveWorldMap(worldMap)
	}
}
