package main

import "os"

type WorldMap [10000]uint8

func CreateWorldMap() WorldMap {
	var world WorldMap
	world[300] = 4
	world[1560] = 4
	world[9780] = 4
	return world //
}

func SaveWorldMap(world WorldMap) error {
	return os.WriteFile("worldmap.dat", world[:], 0644)
}

func LoadWorldMap() (WorldMap, error) {
	var world WorldMap
	data, err := os.ReadFile("worldmap.dat")
	if err != nil {
		return WorldMap{}, err
	}
	copy(world[:], data)
	return world, nil
}
