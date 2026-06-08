package main

type Shape interface {
	getName() string
	getPerimeter() float64
	getSides() (
		float64,
		float64,
		float64,
	)
}

type Triangle struct {
	name string
	a    float64
	b    float64
	c    float64
}

func (t Triangle) getName() string {
	return t.name
}

func (t Triangle) getPerimeter() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) getSides() (float64, float64, float64) {
	return t.a, t.b, t.c
}

func structs_main() {
	var shapes []Shape = []Shape{
		Triangle{name: "Triangle", a: 3, b: 4, c: 5},
	}

	for _, shape := range shapes {
		perimeter := shape.getPerimeter()
		a, b, c := shape.getSides()

		println("------------------")
		println("Shape:", shape.getName())
		println("sides:", a, b, c)
		println("perimeter:", perimeter)
		println("------------------")
	}
}
