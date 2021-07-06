package common_test

import (
	"testing"

	common "github.com/Cyb3r-Jak3/common/go"
)

func TestGoodHash(t *testing.T) {
	hashed, err := common.HashFile("sha1", "../LICENSE")
	if err != nil {
		t.Errorf("Wanted a good hash and got %s", err.Error())
	}
	if hashed != "e2bd6bdb41b8337bf19f0363674998a1319bed52" {
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
	_, err := common.HashFile("", "../LICENSE")
	if err == nil {
		t.Error("Wanted an error but it hashed the file")
	}
}

func TestHashers(t *testing.T) {
	for _, i := range []string{"1", "256", "384", "512"} {
		_, err := common.HashFile(i, "../LICENSE")
		if err != nil {
			t.Errorf("Wanted a good hash with %s and got %s", i, err.Error())
		}
	}
}
