package sim

import (
	"testing"
	"time"

	"github.com/AlexeyNilov/space_sim/space"
	"github.com/stretchr/testify/assert"
)

const spaceSize = 10

func TestCreateParticles(t *testing.T) {
	// Helper function to count points with particles
	countParticles := func(s *space.Space) int {
		count := 0
		for _, point := range *s {
			if point.PointsTo != nil {
				count++
			}
		}
		return count
	}

	tests := []struct {
		name        string
		probability float64
		expected    int
	}{
		{"AllParticles", 1, spaceSize}, // Probability = 1, expect all points to have particles
		{"NoParticles", 0, 0},          // Probability = 0, expect no points to have particles
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := space.CreateSpace(spaceSize)
			CreateParticles(s, tt.probability)
			assert.Equal(t, tt.expected, countParticles(s))
		})
	}
}

func BenchmarkRun(b *testing.B) {
	s := space.CreateSpace(spaceSize)
	CreateParticles(s, 0.5)
	b.ResetTimer()
	for range 10 {
		time.Sleep(500 * time.Millisecond)
		MoveParticles(s)
		s.GC()
	}
}
