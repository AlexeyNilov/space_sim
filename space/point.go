package space

import "github.com/AlexeyNilov/space_sim/particle"

type Point struct {
	Next    *Point
	ID      int
	Content string
}

func NewPoint() *Point {
	return &Point{}
}

func (p *Point) CreateParticle(name string) *particle.Particle {
	return particle.NewParticle(name)
}

func LinkPoints(p1, p2 *Point) {
	p1.Next = p2
}
