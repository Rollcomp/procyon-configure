package configure

type WebServerProperties struct {
	Port            uint   `yaml:"port" json:"port" default:"8080"`
	Shutdown        string `yaml:"shutdown" json:"shutdown" default:"immediate"`
	ShutdownTimeout uint   `yaml:"shutdown.timeout" json:"shutdown.timeout" default:"0"`
}

func newServerProperties() *WebServerProperties {
	return &WebServerProperties{}
}

func (properties *WebServerProperties) GetConfigurationPrefix() string {
	return "server"
}
