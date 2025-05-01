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

func (p *Point) UnlinkParticle() {
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
	if p.PointsTo == nil {
		return
	}

	time.Sleep(10 * time.Millisecond)

	if p.Next.PointsTo == nil && p.PointsTo.GetIntent("") == "Move" {
		p.PointsTo.Energy--
		p.Next.PointsTo = p.PointsTo
		p.UnlinkParticle()
		return
	}

	if p.Next.PointsTo != nil && p.PointsTo.GetIntent(p.Next.GetContent()) == "Eat" {
		particle.ExchangeEnergy(p.PointsTo, p.Next.PointsTo)
	}
}

