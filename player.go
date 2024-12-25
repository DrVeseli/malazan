package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

type Player struct {
	Y uint8 `json:"Y"`
	X uint8 `json:"X"`
}

func CreatePlayer() Player {
	rand.Seed(time.Now().UnixNano())
	return Player{
		Y: uint8(rand.Intn(100)),
		X: uint8(rand.Intn(100)),
	}
}

func SavePlayer(player Player) error {
	data, err := json.Marshal(player)
	if err != nil {
		return err
	}
	return os.WriteFile("player.json", data, 0644)
}

func LoadPlayer() (Player, error) {
	var player Player
	data, err := os.ReadFile("player.json")
	if err != nil {
		return Player{}, err
	}
	err = json.Unmarshal(data, &player)
	return player, err
}

func SetPlayerLocation(player Player, world *WorldMap) {
	index := int(player.Y) * int(player.X)
	if index < len(*world) {
		(*world)[index] = 1
	}
}
