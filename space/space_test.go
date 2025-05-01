package space

import (
	"testing"

	"github.com/AlexeyNilov/space_sim/particle"
	"github.com/stretchr/testify/assert"
)

const spaceSize = 10

func TestAddPoint(t *testing.T) {
	point := NewPoint()
	space := Space{}
	space.AddPoint(point)
	assert.Equal(t, point, &space[0])
	assert.Equal(t, 1, len(space))
}

func TestCreateSpace(t *testing.T) {
	space := CreateSpace(spaceSize)
	assert.Equal(t, spaceSize, len((*space)))
	for i := range *space {
		assert.NotNil(t, (*space)[i].Next)
	}

	// test if the last and first points are linked
	assert.Equal(t, &(*space)[0], (*space)[len(*space)-1].Next)

	// test ID
	assert.Equal(t, 0, (*space)[0].ID)
}

func BenchmarkCreateSpace(b *testing.B) {
	CreateSpace(spaceSize)
}

func TestGetPoint(t *testing.T) {
	space := CreateSpace(spaceSize)
	point := space.GetPoint(0)
	assert.Equal(t, 0, point.ID)
}

func TestGetSize(t *testing.T) {
	space := CreateSpace(spaceSize)
	got := space.GetSize()
	assert.Equal(t, spaceSize, got)
}

func TestGC(t *testing.T) {
	space := CreateSpace(spaceSize)
	if (*space)[0].PointsTo == nil {
		(*space)[0].PointsTo = particle.NewParticle("Empty")
	}
	(*space)[0].PointsTo.Energy = 0
	space.GC()
	// test if the first point is still there
	assert.Nil(t, (*space)[0].PointsTo)
}