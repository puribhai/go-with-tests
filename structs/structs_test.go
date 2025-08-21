package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("check perimeter for some width and height", func(t *testing.T) {
		got := Perimeter(4.0, 5.2)
		want := 18.40

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
func TestArea(t *testing.T) {
	t.Run("check area for some width and height", func(t *testing.T) {
		got := Area(4.0, 5.2)
		want := 20.8

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
}
