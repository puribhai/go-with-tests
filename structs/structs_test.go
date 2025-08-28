package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("check perimeter for some width and height", func(t *testing.T) {

		rectangle := Rectangle{4.0, 5.2}
		got := rectangle.Perimeter()
		want := 18.40

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{4.0, 5.2}
		got := rectangle.Area()
		want := 20.8

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
