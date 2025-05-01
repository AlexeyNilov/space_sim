package particle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParticle(t *testing.T) {
	got := NewParticle("Test")
	assert.IsType(t, &Particle{}, got)
}
