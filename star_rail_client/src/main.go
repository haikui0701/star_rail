package main

import (
	"example.com/m/game"
)

//go build -ldflags="-H windowsgui"

func main() {
	game.GetStarRailClient().Start()
}