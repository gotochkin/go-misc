// Copyright 2022 Gleb Otochkin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
/// Methods for types
/// Type point in two dimention coordinates and different methods
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X,
	Y float64
}

func (p Point) Distance(z Point) float64 {
	return math.Hypot(z.X-p.X, z.Y-p.Y)
}

type Path []Point

func (pt Path) Distance() float64 {
	path := 0.0
	for i := range pt {
		if i > 0 {
			path += pt[i-1].Distance(pt[i])
		}
	}
	return path
}

func main() {
	// Points
	p1 := Point{1, 1}
	p2 := Point{6, 1}
	p3 := Point{6, 5}
	p4 := Point{1, 5}

	// Rectangle 5x4
	k := Path{
		p1,
		p2,
		p3,
		p4,
		p1,
	}
	//Perimeter of the Rectanhgle

	fmt.Println(p1.Distance(p2))
	fmt.Println(p2.Distance(p3))
	fmt.Println(p3.Distance(p4))
	fmt.Println(p4.Distance(p1))
	fmt.Println(k.Distance())
}
