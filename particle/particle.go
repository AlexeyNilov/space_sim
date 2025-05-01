package particle

type Particle struct {
	Name string
}

func NewParticle(name string) *Particle {
	return &Particle{Name: name}
}
