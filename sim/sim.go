package sim

import (
	"log"

	"github.com/AlexeyNilov/space_sim/space"
)

func Run(space *space.Space) {
	log.Println("Start sim")
	pointID := 0
	(*space)[pointID].CreateParticle("P")

}
