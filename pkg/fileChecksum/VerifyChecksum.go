package fileChecksum

import (
	"errors"
	"fmt"
	"path"
	"strings"
)

var (
	Verbose = false
)

func VerifyChecksum(filepath string, checksumUrl string, ExtendedChecksum bool) error {

	var (
		err              error
		checksumFile     []byte
		ValidCheckSum    string
		FileCheckSum     string
		comparisonResult int
	)

	checksumFile, err = retrieveFile(checksumUrl) // Download checksum
	if err != nil {
		return err
	}
	ValidCheckSum = string(checksumFile)

	FileCheckSum, err = GenerateSha256FromFile(filepath) // Generate checksum of downloaded file
	if err != nil {
		return err
	}

	if ExtendedChecksum { // `<checksum>  <filename>` and optionally multiple files (one per line)
		lines := strings.Split(ValidCheckSum, "\n")
		for _, line := range lines {
			parts := strings.Split(line, " ")
			if parts[2] == path.Base(filepath) {
				ValidCheckSum = parts[0]
				break // found match
			}
		}
	} // else `<checksum>` only -> Nothing to do
	comparisonResult = strings.Compare(FileCheckSum, ValidCheckSum) // Compare checksums

	if comparisonResult != 0 { // If checksums are not equal
		if Verbose {
			fmt.Println("Downloaded checksum: ", ValidCheckSum)
			fmt.Println("Actual checksum:     ", FileCheckSum)
		}
		return errors.New("sha256 checksum does not match downloaded file")
	}

	return nil
}
