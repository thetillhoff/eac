package apps

import (
	"runtime"

	"github.com/thetillhoff/eac/internal/templates"
	"github.com/thetillhoff/eac/pkg/logs"
)

func IsValidPlatform(potentialPlatform string) bool {
	for _, platform := range templates.GetPlatforms() {
		if platform == potentialPlatform {
			return true
		}
	}
	return potentialPlatform == "all" // 'all' is allowed, too
}

func ValidPlatforms() []string {
	validPlatforms := templates.GetPlatforms()
	// validPlatforms = append(validPlatforms, "all")
	return validPlatforms
}

func ResolvePlatforms(flaggedPlatforms []string) []string {
	platforms := []string{}
	for _, flaggedPlatform := range flaggedPlatforms {
		if IsValidPlatform(flaggedPlatform) && flaggedPlatform != "all" {
			platforms = append(platforms, flaggedPlatform)
		}
		if flaggedPlatform == "all" {
			if len(flaggedPlatforms) > 1 {
				logs.Warn("You set platforms to all, ignoring other specified platforms.")
			}
			return ValidPlatforms()
		}
	}
	if len(platforms) == 0 { // no specific platforms selected, therefore adding only the current one
		platforms = append(platforms, runtime.GOOS)
	}
	return platforms
}
