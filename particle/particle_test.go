package particle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParticle(t *testing.T) {
	got := NewParticle("Test")
	assert.IsType(t, &Particle{}, got)
}

func TestGetIntent(t *testing.T) {
	p := NewParticle("Test")
	got := p.GetIntent()
	assert.Equal(t, "Wait", got)

	p.Energy = 10
	got = p.GetIntent()
	assert.Equal(t, "Move", got)
}

func TestGetDescription(t *testing.T) {
	p := NewParticle("P")
	got := p.GetDescription()
	assert.Equal(t, "P:0", got)

	p.Energy = 10
	got = p.GetDescription()
	assert.Equal(t, "P:10", got)
}