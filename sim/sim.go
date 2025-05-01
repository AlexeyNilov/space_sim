package sim

import (
	"time"

	"github.com/AlexeyNilov/space_sim/space"
)

func Run(space *space.Space) {
	pointID := 0
	spaceSize := len(*space)

	for range spaceSize {
		time.Sleep(3 * time.Second)
		(*space)[pointID].Content = "P"
		if pointID == spaceSize-1 {
			pointID = 0
		} else {
			pointID++
		}
	}
}
