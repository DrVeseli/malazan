package main

import "os"

type WorldMap [10000]uint8

func CreateWorldMap() WorldMap {
    var world WorldMap
    return world // Go automatically initializes arrays with zero values
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
