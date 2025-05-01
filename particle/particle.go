package particle

import (
	"fmt"
	"math/rand"
)

type Particle struct {
	Name   string
	Energy int
}

func NewParticle(name string) *Particle {
	return &Particle{Name: name, Energy: rand.Intn(10) + 1}
}

func (p *Particle) GetIntent(content string) string {
	if p.Energy > 1 && content == ""{
		return "Move"
	}
	return "Eat"
}

func (p *Particle) GetDescription() string {
	if p.Name != "" {
		return fmt.Sprintf("%s:%d", p.Name, p.Energy)
	}
	return ""
}

func ExchangeEnergy(p1 *Particle, p2 *Particle) {
	if p1.Energy > p2.Energy {
		p1.Energy += 1
		p2.Energy -= 1
	}
}
