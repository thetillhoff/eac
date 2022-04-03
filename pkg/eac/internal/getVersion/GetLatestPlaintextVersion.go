package getVersion

import (
	"strings"

	"github.com/thetillhoff/eac/pkg/eac/internal/version"
)

func GetLatestPlaintextVersion(versionProvider version.VersionProvider) (string, error) {
	content, err := retrieveFile(versionProvider.Url)
	if err != nil {
		return "", err
	}

	result := string(content)

	result = strings.TrimPrefix(result, versionProvider.TrimPrefix)
	result = strings.TrimSuffix(result, versionProvider.TrimSuffix)

	return result, nil
}
