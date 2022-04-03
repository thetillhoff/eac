package helpers

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
)

func DownloadFile(url string, destination string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	// Check response code
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New("file could not be retrieved, statuscode " + strconv.Itoa(resp.StatusCode))
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
