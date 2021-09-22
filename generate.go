package common

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// GenerateRandInt securely generate a random int64. The input is the maximun value that the random int can be
func GenerateRandInt(x int) (returnValue int, returnError error) {
	if x <= 0 {
		return 0, errors.New("need a row amount of greater than 0")
	}
	value, returnError := rand.Int(rand.Reader, big.NewInt(int64(x)))
	returnValue = int(value.Int64())
	return
}
