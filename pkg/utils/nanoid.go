package utils

import "github.com/jaevor/go-nanoid"

// DefaultAlphabet is the default alphabet used.
const DefaultAlphabet = "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
const DefaultLength = 16

func GenerateNanoID(length int, prefix ...string) (string, error) {
	if length == 0 {
		length = DefaultLength
	}
	nanoID, err := nanoid.CustomASCII(DefaultAlphabet, length)
	if err != nil {
		return "", err
	}

	if len(prefix) > 0 {
		return prefix[0] + nanoID(), nil
	}

	return nanoID(), nil
}
