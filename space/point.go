package space

import (
	"time"

	"github.com/AlexeyNilov/space_sim/particle"
)

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
	return p.PointsTo.GetDescription()
}

func LinkPoints(p1, p2 *Point) {
	p1.Next = p2
}

func (p *Point) MoveParticle() {
	if p.PointsTo != nil && p.Next.PointsTo == nil {
		time.Sleep(10 * time.Millisecond)
		if p.PointsTo.GetIntent() == "Move" {
			p.PointsTo.Energy -= 1
			p.Next.PointsTo = p.PointsTo
			p.RemoveParticle()
		}
	}
}
