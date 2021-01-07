package configure

type LoggingProperties struct {
	Level    string `yaml:"level" json:"level" default:"TRACE"`
	FileName string `yaml:"file.name" json:"file.name"`
}

func (properties *LoggingProperties) GetConfigurationPrefix() string {
	return "logging"
}
