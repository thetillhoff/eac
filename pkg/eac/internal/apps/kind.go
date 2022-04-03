package apps

import "github.com/thetillhoff/eac/pkg/eac/internal/version"

func init() {
	Apps["kind"] = App{
		Destination: "/usr/local/bin/kind",
		VersionProvider: version.VersionProvider{
			Provider:   version.Json,
			Url:        "https://api.github.com/repos/kubernetes-sigs/kind/releases/latest",
			JsonQuery:  "tag_name",
			TrimPrefix: "v",
		},
		DownloadUrlTemplate: "https://github.com/kubernetes-sigs/kind/releases/download/v{{ .Version }}/kind-{{ .OS }}-{{ .Arch }}",
		Checksum: &Checksum{
			UrlTemplate:      "https://github.com/kubernetes-sigs/kind/releases/download/v{{ .Version }}/kind-{{ .OS }}-{{ .Arch }}.sha256sum",
			ExtendedChecksum: true,
		},
		PermissionAdjustments: []PermissionAdjustment{
			{
				Path: "/kind",
				Mode: 0755,
			},
		},
		PathAddition: "/usr/local/bin",
	}
}
