package particle

import "fmt"

type Particle struct {
	Name   string
	Energy int
}

func NewParticle(name string) *Particle {
	return &Particle{Name: name, Energy: 5}
}

func (p *Particle) GetIntent() string {
	if p.Energy > 1 {
		return "Move"
	}
	return "Wait"
}

func (p *Particle) GetDescription() string {
	if p.Name != "" {
		return fmt.Sprintf("%s:%d", p.Name, p.Energy)
	}
	return ""
}
