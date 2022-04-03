package apps

import "github.com/thetillhoff/eac/pkg/eac/internal/version"

func init() {
	Apps["helm"] = App{
		Destination: "/usr/local/bin/helm",
		VersionProvider: version.VersionProvider{
			Provider:   version.Json,
			Url:        "https://api.github.com/repos/helm/helm/releases/latest",
			JsonQuery:  "tag_name",
			TrimPrefix: "v",
		},
		DownloadUrlTemplate: "https://get.helm.sh/helm-v{{ .Version }}-{{ .OS }}-{{ .Arch }}.tar.gz",
		Checksum: &Checksum{
			UrlTemplate:      "https://get.helm.sh/helm-v{{ .Version }}-{{ .OS }}-{{ .Arch }}.tar.gz.sha256sum",
			ExtendedChecksum: true,
		},
		TarProperties: &TarProperties{
			Path: "/{{ .OS }}-{{ .Arch }}/helm",
		},
		PathAddition: "/usr/local/bin",
	}
}
