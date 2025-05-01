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
	p.Energy = 0
	got := p.GetIntent("")
	assert.Equal(t, "Eat", got)

	p.Energy = 10
	got = p.GetIntent("")
	assert.Equal(t, "Move", got)
}

func TestGetDescription(t *testing.T) {
	p := NewParticle("P")
	p.Energy = 10
	got := p.GetDescription()
	assert.Equal(t, "P:10", got)
}

func TestExchangeEnergy(t *testing.T) {
	first := NewParticle("Test")
	first.Energy = 5
	second := NewParticle("Test")
	second.Energy = 2
	ExchangeEnergy(first, second)
	assert.Equal(t, 6, first.Energy)
	assert.Equal(t, 1, second.Energy)
}