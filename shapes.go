package nqtree

type Shape interface {
	Contains(p *Point) (bool, error)
	Edges() []*Point
}

type Orthotope struct {
	MinPoint Point
	MaxPoint Point
}

func (o Orthotope) Contains(p *Point) (bool, error) {
	for index, min := range o.MinPoint.Dimensions {
		max := o.MaxPoint.Dimensions[index]
		if p.Dimensions[index] < min || p.Dimensions[index] > max {
			return false, nil
		}
	}
	return true, nil
}

func (o Orthotope) Edges() []*Point {
	edges := []*Point{&o.MinPoint, &o.MinPoint}
	return edges
}

type HyperSphere struct {
	CenterPoint Point
	Radius      float64
}

func (s HyperSphere) Contains(p *Point) (bool, error) {
	distance, err := s.CenterPoint.Distance(p)
	if err != nil {
		return false, err
	}
	return distance < s.Radius, nil
}

func (s HyperSphere) Edges() []*Point {
	edges := []*Point{&s.CenterPoint}
	return edges
}
