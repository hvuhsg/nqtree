package nqtree

type TreeNode struct {
	Rect       Orthotope
	Childs     []TreeNode
	DataPoints []DataPoint
	MaxPoints  int
	Divided    bool
}

func (t *TreeNode) Insert(dp *DataPoint) (bool, error) {
	contains, err := t.Rect.Contains(&dp.Position)
	if err != nil {
		return false, err
	}
	if !contains {
		return false, nil
	}
	if len(t.DataPoints) < t.MaxPoints {
		t.DataPoints = append(t.DataPoints, *dp)
		return true, nil
	}
	if !t.Divided {
		t.Divide()
		t.Divided = true
	}
	for i := 0; i < len(t.Childs); i++ {
		child := &t.Childs[i]
		inserted, err := child.Insert(dp)
		if err != nil {
			return false, err
		}
		if inserted {
			return true, nil
		}
	}
	return false, nil
}

func (t *TreeNode) Divide() {
	var childs []Orthotope
	childs2 := []Orthotope{t.Rect}
	for dimIndex := 0; dimIndex < len(t.Rect.MinPoint.Dimensions); dimIndex++ {
		childs = childs2
		childs2 = []Orthotope{}
		for _, c := range childs {
			middle := (c.MinPoint.Dimensions[dimIndex] + c.MaxPoint.Dimensions[dimIndex]) / 2
			var dimensionsMax []float64 = make([]float64, len(c.MaxPoint.Dimensions))
			copy(dimensionsMax, c.MaxPoint.Dimensions)
			dimensionsMax[dimIndex] = middle
			var dimensionsMin []float64 = make([]float64, len(c.MinPoint.Dimensions))
			copy(dimensionsMin, c.MinPoint.Dimensions)
			dimensionsMin[dimIndex] = middle
			var middlePointMax Point = Point{dimensionsMax}
			var middlePointMin Point = Point{dimensionsMin}
			childs2 = append(
				childs2,
				Orthotope{MinPoint: c.MinPoint, MaxPoint: middlePointMax},
				Orthotope{MinPoint: middlePointMin, MaxPoint: c.MaxPoint},
			)
		}
	}
	for _, c := range childs2 {
		t.Childs = append(t.Childs, TreeNode{Rect: c, MaxPoints: t.MaxPoints, Divided: false})
	}
}

func (t *TreeNode) intresect(s Shape) bool {
	for _, edge := range s.Edges() {
		contained, err := t.Rect.Contains(edge)
		if err != nil {
			return false
		}
		if contained {
			return true
		}
	}
	for _, edge := range t.Rect.Edges() {
		contained, err := s.Contains(edge)
		if err != nil {
			return false
		}
		if contained {
			return true
		}
	}
	return false
}

func (t *TreeNode) containedDataPoints(s Shape) ([]*DataPoint, error) {
	points := []*DataPoint{}
	for index := range t.DataPoints {
		dataPoint := &t.DataPoints[index]
		contained, err := s.Contains(&dataPoint.Position)
		if err != nil {
			return nil, err
		}
		if contained {
			points = append(points, dataPoint)
		}
	}
	return points, nil
}

func (t *TreeNode) Search(s Shape) ([]*DataPoint, error) {
	var dataPoints []*DataPoint
	if t.intresect(s) {
		containedPoints, err := t.containedDataPoints(s)
		if err != nil {
			return nil, err
		}
		dataPoints = append(dataPoints, containedPoints...)
		for _, child := range t.Childs {
			points, err := child.Search(s)
			if err != nil {
				return nil, err
			}
			dataPoints = append(dataPoints, points...)
		}
	} else {
		return dataPoints, nil
	}
	return dataPoints, nil
}
