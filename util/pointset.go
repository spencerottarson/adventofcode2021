package util

type Point struct {
	Row int
	Col int
	Value int
}

type PointSet struct {
	points map[*Point]bool
}

func NewPointSet() PointSet {
	return PointSet{make(map[*Point]bool)}
}

func (s *PointSet) Add(point *Point) {
	s.points[point] = true
}

func (s *PointSet) Remove(point *Point) {
	s.points[point] = false
}

func (s *PointSet) Pop() *Point {
	for point, value := range s.points {
		if value == true {
			s.points[point] = false
			return point
		}
	}

	return nil
}

func (s *PointSet) Contains(point *Point) bool {
	return s.points[point] == true
}

func (s *PointSet) Size() int {
	count := 0
	for _, value := range s.points {
		if value == true {
			count++
		}
	}

	return count
}

func (s *PointSet) IsEmpty() bool {
	for _, value := range s.points {
		if value == true {
			return false
		}
	}

	return true
}