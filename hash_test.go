package common_test

import (
	"testing"

	common "github.com/Cyb3r-Jak3/common/v2"
)

func TestGoodHash(t *testing.T) {
	hashed, err := common.HashFile("256", "./LICENSE")
	if err != nil {
		t.Errorf("Wanted a good hash and got %s", err.Error())
	}
	if hashed != "1f256ecad192880510e84ad60474eab7589218784b9a50bc7ceee34c2b91f1d5" {
		t.Errorf("Mismatched hash. Got %s", hashed)
	}
}

func TestMissingHash(t *testing.T) {
	_, err := common.HashFile("", "missing")
	if err == nil {
		t.Error("Wanted an error but it hashed the file")
	}
}

func TestUnknownHash(t *testing.T) {
	_, err := common.HashFile("", "./LICENSE")
	if err == nil {
		t.Error("Wanted an error but it hashed the file")
	}
}

func TestHashers(t *testing.T) {
	for _, i := range []string{"256", "384", "512"} {
		_, err := common.HashFile(i, "./LICENSE")
		if err != nil {
			t.Errorf("Wanted a good hash with %s and got %s", i, err.Error())
		}
	}
}
