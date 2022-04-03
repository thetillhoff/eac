package apps

import "github.com/thetillhoff/eac/pkg/eac/internal/version"

func init() {
	Apps["golang"] = App{
		Destination: "/usr/local/go",
		VersionProvider: version.VersionProvider{
			Provider:   version.Json,
			Url:        "https://go.dev/dl/?mode=json",
			JsonQuery:  "[0].version",
			TrimPrefix: "go",
		},
		DownloadUrlTemplate: "https://go.dev/dl/go{{ .Version }}.{{ .OS }}-{{ .Arch }}.tar.gz",
		// TODO add checksum
		TarProperties: &TarProperties{
			Path: "/go/",
		},
		PathAddition: "/usr/local/go/bin",
		ProfileVariables: []EnvironmentVariables{
			{
				Key:   "GOROOT",
				Value: "/usr/local/go",
			},
			{
				Key:   "GOPATH",
				Value: "$HOME/go",
			},
		},
	}
}
