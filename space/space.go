package space

type Space []Point


func (s *Space) AddPoint(p *Point) {
	*s = append(*s, *p)
}

func CreateSpace(n int) *Space {
	space := make(Space, n) // Pre-allocate space for n points
	for i := range n {
		space[i] = *NewPoint() // Initialize each point in the space
		space[i].ID = i
		LinkPoints(&space[i], &space[(i+1)%n])
	}

	return &space
}

func (s *Space) GetPoint(id int) *Point {
	for i := range *s {
		if (*s)[i].ID == id {
			return &(*s)[i]
		}
	}
	return nil
}
