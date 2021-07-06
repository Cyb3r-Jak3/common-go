package common

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"os"
)

func HashFile(algorithm string, filename string) (value string, err error) {
	f, err := os.Open(filename) // #nosec
	if err != nil {
		err = fmt.Errorf("couldn't open %s. error reason %s", filename, err.Error())
		return
	}
	var hasher hash.Hash
	switch algorithm {
	case "sha1", "1":
		hasher = sha1.New()
	case "sha256", "256":
		hasher = sha256.New()
	case "sha384", "384":
		hasher = sha512.New384()
	case "sha512", "512":
		hasher = sha512.New()
	default:
		err = fmt.Errorf("unsupported algorithm %s ", algorithm)
		return
	}
	if _, err = io.Copy(hasher, f); err != nil {
		err = fmt.Errorf("couldn't hash %s. error reason %s", filename, err.Error())
		return
	}
	if err = f.Close(); err != nil {
		err = fmt.Errorf("error closing file: %s", err.Error())
		return
	}
	value = string(hasher.Sum(nil))
	return
}
