package configure

import (
	context "github.com/Rollcomp/procyon-context"
	"github.com/codnect/go-one"
)

type ServerProperties struct {
	Port              int `yaml:"port" json:"port" default:"8080"`
	ConnectionTimeout int `yaml:"connection-timeout" json:"connection-timeout"`
}

func (properties ServerProperties) GetPrefix() string {
	return "server"
}

type ServerConfiguration struct {
}

func NewServerAutoConfiguration() ServerConfiguration {
	return ServerConfiguration{}
}

func (config ServerConfiguration) Register() []one.Func {
	return nil
}

func (config ServerConfiguration) ConfigurationProperties() []context.ConfigurationProperties {
	return []context.ConfigurationProperties{
		&ServerProperties{},
	}
}
