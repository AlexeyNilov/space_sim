package html

import (
	"bytes"
	"fmt"
	"html/template"
	"math"

	"github.com/AlexeyNilov/space_sim/space"
)

const divClass = `<div class="cell" id="point-%d">%s</div>`

type BoardData struct {
	Side int
	Grid template.HTML // Use template.HTML to indicate that Grid contains HTML
}

func getPointData(space *space.Space, pointIndex int) string {
	content := (*space)[pointIndex].GetContent()
	return fmt.Sprintf(divClass, pointIndex, content)
}

func GenerateGrid(space *space.Space) (int, string) {
	spaceSize := space.GetSize()
	side := int(math.Ceil(math.Sqrt(float64(spaceSize))))
	if side%2 == 0 {
		side++ // Ensure the side length is odd for symmetry
	}

	// Create a grid with empty cells
	grid := make([][]string, side)
	for i := range grid {
		grid[i] = make([]string, side)
		for j := range grid[i] {
			grid[i][j] = `<div class="cell empty"></div>`
		}
	}

	pointIndex := 0

	// Fill the top row
	for col := 0; col < side && pointIndex < spaceSize; col++ {
		grid[0][col] = getPointData(space, pointIndex)
		pointIndex++
	}

	// Fill the right column
	for row := 1; row < side-1 && pointIndex < spaceSize; row++ {
		grid[row][side-1] = getPointData(space, pointIndex)
		pointIndex++
	}

	// Fill the bottom row
	for col := side - 1; col >= 0 && pointIndex < spaceSize; col-- {
		grid[side-1][col] = getPointData(space, pointIndex)
		pointIndex++
	}

	// Fill the left column
	for row := side - 2; row > 0 && pointIndex < spaceSize; row-- {
		grid[row][0] = getPointData(space, pointIndex)
		pointIndex++
	}

	// Flatten the grid to a single string
	var buffer bytes.Buffer
	for _, row := range grid {
		for _, cell := range row {
			buffer.WriteString(cell + "\n")
		}
	}

	// Return the grid as HTML
	return side, buffer.String()
}
