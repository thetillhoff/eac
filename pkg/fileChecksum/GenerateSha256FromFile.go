package fileChecksum

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"io"
	"os"
)

func GenerateSha256FromFile(filepath string) (string, error) {
	var (
		err    error
		file   *os.File
		hasher hash.Hash
	)

	file, err = os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher = sha256.New()
	if _, err = io.Copy(hasher, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), err
}
