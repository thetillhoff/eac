package config

type Config struct {
	ContinueOnError bool     `mapstructure:"continue-on-error"`
	Platforms       []string `mapstructure:"platform"`
	Verbose         bool     `mapstructure:"verbose"`

	UserHomeDir      string
	EacDirPath       string
	AppsDirPath      string `mapstructure:"apps-dir"`
	VersionsFilePath string `mapstructure:"versions-file"`

	ConfigureConfig `mapstructure:",squash"`
	CreateConfig    `mapstructure:",squash"`
	DeleteConfig    `mapstructure:",squash"`
	InitConfig      `mapstructure:",squash"`
	InstallConfig   `mapstructure:",squash"`
	ListConfig      `mapstructure:",squash"`
	UninstallConfig `mapstructure:",squash"`
	UpdateConfig    `mapstructure:",squash"`
	ValidateConfig  `mapstructure:",squash"`
}

type ConfigureConfig struct {
}

type CreateConfig struct {
	GitHubUser string `mapstructure:"github-user"`
}

type DeleteConfig struct {
}

type InitConfig struct {
}

type InstallConfig struct {
	NoConfiguration bool `mapstructure:"no-config"`
	Update          bool `mapstructure:"update"`
	Latest          bool `mapstructure:"latest"`
	//TODO If Update and Latest are combined, it should be updated automatically to the latest version.
}

type ListConfig struct {
	NoVersion bool   `mapstructure:"no-version"`
	Seperator string `mapstructure:"seperator"`
}

type UninstallConfig struct {
}

type UpdateConfig struct {
	DryRun    bool `mapstructure:"dry-run"`
	SkipLocal bool `mapstructure:"skip-local"`
}

type ValidateConfig struct {
}
