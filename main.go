package main

import (
	"flag"

	"github.com/AlexeyNilov/space_sim/html"
	"github.com/AlexeyNilov/space_sim/sim"
	"github.com/AlexeyNilov/space_sim/space"
)

func main() {
	// Parse flags
	var addr = flag.String("addr", "127.0.0.1:8080", "The address:port combination to listen on for HTTP requests.")
	flag.Parse()

	const spaceSize = 16 // Total number of points
	space := space.CreateSpace(spaceSize)
	side, grid := html.GenerateGrid(space)

	go sim.Run(space)
	
	html.StartServer(addr, side, grid, space)
}
