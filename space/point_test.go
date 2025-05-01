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
	got = point.GetContent()
	assert.Equal(t, "Test", got)

	point.CreateParticle("")
	got = point.GetContent()
	assert.Equal(t, "", got)
}

func TestRemoveParticle(t *testing.T) {
	point := NewPoint()
	point.CreateParticle("Test")
	point.RemoveParticle()

	assert.Nil(t, point.PointsTo)
}

