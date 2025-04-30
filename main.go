package main

import (
	"github.com/AlexeyNilov/space_sim/html"
	"github.com/AlexeyNilov/space_sim/space"
)

func main() {
	const spaceSize = 16 // Total number of points

	space := space.CreateSpace(spaceSize)
	side, grid := html.GenerateGrid(space)
	html.GenerateHTML(side, grid)
}
