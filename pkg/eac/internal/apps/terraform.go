package apps

import "github.com/thetillhoff/eac/pkg/eac/internal/version"

func init() {
	Apps["terraform"] = App{
		Destination: "/usr/local/bin/terraform",
		VersionProvider: version.VersionProvider{
			Provider:   version.Json,
			Url:        "https://api.github.com/repos/hashicorp/terraform/releases/latest",
			JsonQuery:  "tag_name",
			TrimPrefix: "v",
		},
		DownloadUrlTemplate: "https://releases.hashicorp.com/terraform/{{ .Version }}/terraform_{{ .Version }}_{{ .OS }}_{{ .Arch }}.zip",
		Checksum: &Checksum{
			UrlTemplate:      "https://releases.hashicorp.com/terraform/{{ .Version }}/terraform_{{ .Version }}_SHA256SUMS",
			ExtendedChecksum: true,
		},
		ZipProperties: &ZipProperties{
			Path: "/terraform",
		},
		PathAddition: "/usr/local/bin",
	}
}
