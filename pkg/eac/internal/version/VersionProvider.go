package version

type VersionProvider struct {
	Provider   ProviderType
	Url        string
	JsonQuery  string // only for providertype Json
	TrimPrefix string
	TrimSuffix string
}
