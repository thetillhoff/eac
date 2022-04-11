package eac

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/thetillhoff/eac/pkg/eac/internal/apps"
	"github.com/thetillhoff/eac/pkg/eac/internal/helpers"
	"github.com/thetillhoff/eac/pkg/fileChecksum"
	"github.com/thetillhoff/eac/pkg/stringTemplate"
	"github.com/thetillhoff/eac/pkg/untar"
	"github.com/thetillhoff/eac/pkg/unzip"
)

// Download specific version of specific app.
// In case of .tar, .tar.gz or .zip files, they are automatically extracted and the archive deleted.
// Returns filepath of the download or folderpath depending on whether extraction was involved or not.
func DownloadAppVersion(appName string, version string) (string, error) {
	var (
		err          error
		app          = apps.Apps[appName]
		mappedValues = map[string]interface{}{}
		downloadUrl  string
		checksumUrl  string
		tarPath      string
		zipPath      string
	)

	mappedValues["OS"] = runtime.GOOS
	mappedValues["Arch"] = runtime.GOARCH
	mappedValues["Version"] = version

	downloadUrl, err = stringTemplate.StringTemplate(app.DownloadUrlTemplate, mappedValues)
	if err != nil {
		return "", err
	}

	filepath := path.Join(workdir, appName, version, path.Base(downloadUrl))

	if _, err := os.Stat(path.Dir(filepath)); os.IsNotExist(err) { // Only if version was not already downloaded

		if Verbose {
			fmt.Println("Downloading", downloadUrl, "...")
		}

		if !DryRun { // Not if dry-running
			os.MkdirAll(path.Dir(filepath), os.ModePerm) // Creating containing folder

			err = helpers.DownloadFile(downloadUrl, filepath)
			if err != nil {
				return "", err
			}
		}

		if Verbose {
			fmt.Println("Download finished.")
		}

		if app.Checksum != nil && app.Checksum.UrlTemplate != "" { // If checksumUrlTemplate is set, verify it
			if Verbose {
				fmt.Println("Verifying checksum...")
			}

			checksumUrl, err = stringTemplate.StringTemplate(app.Checksum.UrlTemplate, mappedValues) // Generate checksum url
			if err != nil {
				return "", err
			}

			if !DryRun { // Not if dry-running
				fileChecksum.Verbose = Verbose
				err = fileChecksum.VerifyChecksum(filepath, checksumUrl, app.Checksum.ExtendedChecksum)
				if err != nil {
					return "", err
				}
			}

			if Verbose {
				fmt.Println("Checksum verified.")
			}
		}

		if app.TarProperties != nil {

			if Verbose {
				fmt.Println("Extracting...")
			}

			if !DryRun { // Not if dry-running
				untar.Verbose = Verbose
				err = untar.Untar(filepath)
				if err != nil {
					return "", err
				}
			}

			if Verbose {
				fmt.Println("Extraction finished.")
				fmt.Println("Removing archive...")
			}

			if !DryRun { // Not if dry-running
				err = os.Remove(filepath) // Remove tar file
				if err != nil {
					return "", err
				}
			}

			if Verbose {
				fmt.Println("Archive removed.")
			}
		}

		if app.ZipProperties != nil {

			if Verbose {
				fmt.Println("Extracting...")
			}

			if !DryRun { // Not if dry-running
				unzip.Verbose = Verbose
				err = unzip.Unzip(filepath)
				if err != nil {
					return "", err
				}
			}

			if Verbose {
				fmt.Println("Extraction finished.")
				fmt.Println("Removing archive...")
			}

			if !DryRun { // Not if dry-running
				err = os.Remove(filepath) // Remove zip file
				if err != nil {
					return "", err
				}
			}

			if Verbose {
				fmt.Println("Archive removed.")
			}
		}
	}

	if app.TarProperties != nil {
		tarPath, err = stringTemplate.StringTemplate(app.TarProperties.Path, mappedValues)
		if err != nil {
			return "", err
		}

		filepath = path.Join(path.Dir(filepath), tarPath)
	} else if app.ZipProperties != nil {
		zipPath, err = stringTemplate.StringTemplate(app.ZipProperties.Path, mappedValues)
		if err != nil {
			return "", err
		}

		filepath = path.Join(path.Dir(filepath), zipPath)
	}

	return filepath, nil
}
