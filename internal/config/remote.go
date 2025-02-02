package config

type Remote struct {
	GitURL string `json:"git_url,omitempty" koanf:"git_url"              mapstructure:"git_url" toml:"git_url"    yaml:"git_url"`
	Ref    string `json:"ref,omitempty"     mapstructure:"ref,omitempty" toml:"ref,omitempty"   yaml:",omitempty"`
	// Deprecated
	Config           string   `json:"config,omitempty"            mapstructure:"config,omitempty"  toml:"config,omitempty"                    yaml:",omitempty"`
	Configs          []string `json:"configs,omitempty"           mapstructure:"configs,omitempty" toml:"configs,omitempty"                   yaml:",omitempty"`
	Refetch          bool     `json:"refetch,omitempty"           mapstructure:"refetch,omitempty" toml:"refetch,omitempty"                   yaml:",omitempty"`
	RefetchFrequency string   `json:"refetch_frequency,omitempty" koanf:"refetch_frequency"        mapstructure:"refetch_frequency,omitempty" toml:"refetch_frequency,omitempty" yaml:",omitempty"`
}

func (r *Remote) Configured() bool {
	if r == nil {
		return false
	}

	return len(r.GitURL) > 0
}
