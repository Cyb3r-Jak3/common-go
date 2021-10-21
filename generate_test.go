package common_test

import (
	"testing"

	"github.com/Cyb3r-Jak3/common/v3"
)

func TestGenerate(t *testing.T) {
	_, err := common.GenerateRandInt(5)
	if err != nil {
		t.Errorf("Wanted no error. Got %s", err)
	}
	result, err := common.GenerateRandInt(-1)
	if err == nil {
		t.Error("Wanted an error and didn't get on")
	}
	if result != 0 {
		t.Errorf("Result needed to be 0 and got %d", result)
	}
}
