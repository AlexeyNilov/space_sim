package space

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPoint(t *testing.T) {
	got := NewPoint()
	assert.IsType(t, &Point{}, got)
}

func TestLinkPoints(t *testing.T) {
	point1 := NewPoint()
	point2 := NewPoint()
	LinkPoints(point1, point2)
	assert.Equal(t, point2, point1.Next)
}

func TestCreateParticle(t *testing.T) {
	point := NewPoint()
	point.CreateParticle("Test")

	assert.Equal(t, "Test", point.PointsTo.Name)
}

func TestGetContent(t *testing.T) {
	point := NewPoint()

	got := point.GetContent()
	assert.Equal(t, "", got)

	point.CreateParticle("Test")
	point.PointsTo.Energy = 0
	got = point.GetContent()
	assert.Equal(t, "Test:0", got)

	point.CreateParticle("")
	got = point.GetContent()
	assert.Equal(t, "", got)
}

func TestUnlinkParticle(t *testing.T) {
	point := NewPoint()
	point.CreateParticle("Test")
	point.UnlinkParticle()

	assert.Nil(t, point.PointsTo)
}

func TestMoveParticle(t *testing.T) {
	startPoint := NewPoint()
	startPoint.CreateParticle("First")
	endPoint := NewPoint()
	startPoint.Next = endPoint
	startPoint.MoveParticle()

	assert.Nil(t, startPoint.PointsTo)
	assert.Equal(t, "First", endPoint.PointsTo.Name)

	startPoint.MoveParticle()
	assert.Equal(t, "First", endPoint.PointsTo.Name)

	startPoint.CreateParticle("Second")
	startPoint.MoveParticle()
	assert.Equal(t, "Second", startPoint.PointsTo.Name)
	assert.Equal(t, "First", endPoint.PointsTo.Name)
}
