package configure

type WebServerProperties struct {
	Port              int `yaml:"port" json:"port" default:"8080"`
	ConnectionTimeout int `yaml:"connection-timeout" json:"connection-timeout"`
}

func newServerProperties() *WebServerProperties {
	return &WebServerProperties{}
}

func (properties *WebServerProperties) GetPrefix() string {
	return "server"
}
