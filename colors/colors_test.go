package colors

import (
	"testing"

)

func TestColors(t *testing.T) {
	if Aqua != 1752220 {
		t.Error("Aqua color is wrong")
	}
	hexColor := ToHex(Aqua)
	if hexColor != "1ABC9C" {
		t.Errorf("Hex color should '1ABC9C' be and got %s", hexColor)
	}
}


func BenchmarkToHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToHex(Aqua)
	}
}