package apps

import "github.com/thetillhoff/eac/pkg/eac/internal/version"

func init() {
	Apps["github-cli"] = App{
		Destination: "/usr/local/bin/eac",
		VersionProvider: version.VersionProvider{
			Provider:   version.Json,
			Url:        "https://api.github.com/repos/thetillhoff/eac/releases/latest",
			JsonQuery:  "tag_name",
			TrimPrefix: "v",
		},
		DownloadUrlTemplate: "https://github.com/thetillhoff/eac/releases/download/v{{ .Version }}/eac_{{ .OS }}_{{ .Arch }}",
		// TODO add checksum
		PermissionAdjustments: []PermissionAdjustment{
			{
				Path: "/eac",
				Mode: 0755,
			},
		},
		PathAddition: "/usr/local/bin",
	}
}
