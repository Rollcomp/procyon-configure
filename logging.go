package configure

type LoggingProperties struct {
	Level string `yaml:"level" json:"level" default:"TRACE"`
	File  string `yaml:"file" json:"file"`
	Path  string `yaml:"path" json:"path"`
}

func newLoggingProperties() *LoggingProperties {
	return &LoggingProperties{}
}

func (properties *LoggingProperties) GetConfigurationPrefix() string {
	return "logging"
}
