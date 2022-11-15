package common

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	_, err := GenerateRandInt(5)
	if err != nil {
		t.Errorf("Wanted no error. Got %s", err)
	}
	result, err := GenerateRandInt(-1)
	if err == nil {
		t.Error("Wanted an error and didn't get on")
	}
	if result != 0 {
		t.Errorf("Result needed to be 0 and got %d", result)
	}
}
