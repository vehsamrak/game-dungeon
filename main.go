package main

import "github.com/vehsamrak/game-dungeon/internal/client/console"

func main() {
	client := console.Client{}.Create()
	client.Start()
}
