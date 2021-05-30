package apps

import (
	"runtime"

	"github.com/thetillhoff/eac/pkg/logs"
)

func IsValidPlatform(potentialPlatform string) bool {
	for platform := range demoFiles {
		if platform == potentialPlatform {
			return true
		}
	}
	return potentialPlatform == "all" // 'all' is allowed, too
}

func ValidPlatforms() []string {
	validPlatforms := make([]string, 0, len(demoFiles))
	for platform := range demoFiles {
		validPlatforms = append(validPlatforms, platform)
	}
	validPlatforms = append(validPlatforms, "all")
	return validPlatforms
}

func ResolvePlatforms(flaggedPlatforms []string) []string {
	platforms := []string{}
	for _, flaggedPlatform := range flaggedPlatforms {
		if IsValidPlatform(flaggedPlatform) {
			platforms = append(platforms, flaggedPlatform)
		}
		if flaggedPlatform == "all" {
			if len(platforms) > 0 {
				logs.Warn("You set platfroms to all, ignoring other specified platforms.")
			}
			platforms = ValidPlatforms()
			continue
		}
	}
	if len(platforms) == 0 { // no specific platforms selected
		platforms = append(platforms, runtime.GOOS)
	}
	return platforms
}
