package getVersion

import (
	"errors"
	"strings"

	"github.com/elgs/gojq"
	"github.com/thetillhoff/eac/pkg/eac/internal/version"
)

func GetLatestJsonVersion(versionProvider version.VersionProvider) (string, error) {
	var (
		err  error
		body []byte
	)

	body, err = retrieveFile(versionProvider.Url)
	if err != nil {
		return "", err
	}

	parser, err := gojq.NewStringQuery(string(body))
	if err != nil {
		return "", err
	}

	result, err := parser.Query(versionProvider.JsonQuery)
	if err != nil {
		return "", err
	}

	stringResult, ok := result.(string)
	if !ok {
		return "", errors.New("could not query json: bad type")
	}

	stringResult = strings.TrimPrefix(stringResult, versionProvider.TrimPrefix)
	stringResult = strings.TrimSuffix(stringResult, versionProvider.TrimSuffix)

	return stringResult, err

}
