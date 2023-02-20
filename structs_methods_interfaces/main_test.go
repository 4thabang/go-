package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("expected: %.2f, got: %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	testCases := []struct {
		name       string
		shape      Shaper
		wantedArea float64
	}{
		{name: "rectangles", shape: Rectangle{10.0, 10.0}, wantedArea: 100.0},
		{name: "circles", shape: Circle{10}, wantedArea: 314.1592653589793},
		{name: "triangles", shape: Triangle{12, 6}, wantedArea: 36.0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			CheckArea(t, tt.shape, tt.wantedArea)
		})
	}
}

func CheckArea(t testing.TB, shape Shaper, want float64) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("%#v - expected: %g, got: %g", shape, want, got)
	}
}
