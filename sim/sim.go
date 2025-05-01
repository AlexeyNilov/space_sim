package sim

import (
	"log"
	"math/rand/v2"
	"time"

	"github.com/AlexeyNilov/space_sim/space"
)

func Run(space *space.Space) {
	log.Println("Start sim")
	time.Sleep(1 * time.Second)
	CreateParticles(space, 0.2)
}

func CreateParticles(space *space.Space, probability float64) {
	// Iterate over each point in the space
	for i := range *space {
		// Generate a random number between 0 and 1
		if rand.Float64() < probability {
			// Create a particle at the point if the random number is less than the probability
			(*space)[i].CreateParticle("P")
		}
	}
}
