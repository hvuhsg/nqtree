# NQTree
N Dimensions quadtree lib


#### Installation
```
$> go get github.com/hvuhsg/nqtree
```

#### Usage

```go
package main

import (
	"fmt"
	"nqtree/github.com/hvuhsg/nqtree"
)

func main() {
	dp1 := nqtree.DataPoint{
		Position: nqtree.Point{Dimensions: []float64{50, 50}},
		Data: "In-HyperSphere-and-Orthotope",
	}
	dp2 := nqtree.DataPoint{
		Position: nqtree.Point{Dimensions: []float64{75, 75}},
		Data: "In-Orthotope",
	}
	dp3 := nqtree.DataPoint{
		Position: nqtree.Point{Dimensions: []float64{25, 25}},
		Data: "In-HyperSphere",
	}

	tree := nqtree.TreeNode{
		Rect: nqtree.Orthotope{
			MinPoint: nqtree.Point{Dimensions: []float64{0, 0}},
			MaxPoint: nqtree.Point{Dimensions: []float64{100, 100}},
		},
		MaxPoints: 2,
	}
	tree.Insert(&dp1)
	tree.Insert(&dp2)
	tree.Insert(&dp3)

	s := nqtree.HyperSphere{
		CenterPoint: nqtree.Point{Dimensions: []float64{30, 30}},
		Radius:      30.,
	}
	o := nqtree.Orthotope{
		MinPoint: nqtree.Point{Dimensions: []float64{26, 26}},
		MaxPoint: nqtree.Point{Dimensions: []float64{78, 78}},
	}

	fmt.Printf("\nHyperSphere Search: %v\n", s)
	searchPoints, err := tree.Search(s)
	if err != nil {
		fmt.Println(err)
	}
	for _, p := range searchPoints {
		fmt.Println("	", *p)
	}
	fmt.Printf("\nOrthotope Search: %v\n", o)
	searchPoints, err = tree.Search(o)
	if err != nil {
		fmt.Println(err)
	}
	for _, p := range searchPoints {
		fmt.Println("	", *p)
	}
}
```
