package configure

import (
	"github.com/codnect/go-one"
	context "github.com/procyon-projects/procyon-context"
)

type DataSourceProperties struct {
	DriverName   string `yaml:"driver-name" json:"driver-name"`
	Username     string `yaml:"username" json:"username"`
	Password     string `yaml:"password" json:"password"`
	DatabaseName string `yaml:"database-name" json:"database-name"`
}

func (properties *DataSourceProperties) GetPrefix() string {
	return "procyon.datasource"
}

type DataSourceConfiguration struct {
}

func NewDataSourceConfiguration() DataSourceConfiguration {
	return DataSourceConfiguration{}
}

func (config DataSourceConfiguration) Register() []one.Func {
	return nil
}

func (config DataSourceConfiguration) ConfigurationProperties() []context.ConfigurationProperties {
	return []context.ConfigurationProperties{
		&DataSourceProperties{},
	}
}
