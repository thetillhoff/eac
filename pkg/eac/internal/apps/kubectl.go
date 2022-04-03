package apps

import "github.com/thetillhoff/eac/pkg/eac/internal/version"

func init() {
	Apps["kubectl"] = App{
		Destination: "/usr/local/bin/kubectl",
		VersionProvider: version.VersionProvider{
			Provider:   version.Plaintext,
			Url:        "https://storage.googleapis.com/kubernetes-release/release/stable.txt",
			TrimPrefix: "v",
		},
		DownloadUrlTemplate: "https://dl.k8s.io/release/v{{ .Version }}/bin/{{ .OS }}/{{ .Arch }}/kubectl",
		Checksum: &Checksum{
			UrlTemplate: "https://dl.k8s.io/v{{ .Version }}/bin/{{ .OS }}/{{ .Arch }}/kubectl.sha256",
		},
		PermissionAdjustments: []PermissionAdjustment{
			{
				Path: "/kubectl",
				Mode: 0755,
			},
		},
		PathAddition: "/usr/local/bin",
	}
}
