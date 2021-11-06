package colors_test

import (
	"testing"

	"github.com/Cyb3r-Jak3/common/v4/colors"
)

func TestColors(t *testing.T) {
	if colors.Aqua != 1752220 {
		t.Error("Aqua color is wrong")
	}
	hexColor := colors.ToHex(colors.Aqua)
	if hexColor != "1ABC9C" {
		t.Errorf("Hex color should '1ABC9C' be and got %s", hexColor)
	}
}
