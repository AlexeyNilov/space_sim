package space

import "github.com/AlexeyNilov/space_sim/particle"

type Point struct {
	Next     *Point
	ID       int
	PointsTo *particle.Particle
}

func NewPoint() *Point {
	return &Point{}
}

func (p *Point) CreateParticle(name string) {
	p.PointsTo = particle.NewParticle(name)
}

func (p *Point) RemoveParticle() {
	p.PointsTo = nil
}

func (p *Point) GetContent() string {
	if p.PointsTo == nil {
		return ""
	}
	return p.PointsTo.Name
}

func LinkPoints(p1, p2 *Point) {
	p1.Next = p2
}
