package common

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"os"
)

// HashFile generates a string hash of the given file path. Supported hashing algorithm: sha256, sha384, and sha512.
func HashFile(algorithm string, filepath string) (value string, err error) {
	f, err := os.Open(filepath) // #nosec
	if err != nil {
		err = fmt.Errorf("couldn't open %s: %w", filepath, err)
		return
	}
	var hasher hash.Hash
	switch algorithm {
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
		err = fmt.Errorf("couldn't hash %s: %w", filepath, err)
		return
	}
	if err = f.Close(); err != nil {
		err = fmt.Errorf("error closing file: %w", err)
		return
	}
	value = fmt.Sprintf("%x", hasher.Sum(nil))
	return
}
