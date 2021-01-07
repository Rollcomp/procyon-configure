package configure

type LoggingProperties struct {
	Level    string `yaml:"level" json:"level" default:"DEBUG"`
	FileName string `yaml:"file.name" json:"file.name"`
	FilePath string `yaml:"file.path" json:"file.path"`
}

func (properties *LoggingProperties) GetConfigurationPrefix() string {
	return "logging"
}
