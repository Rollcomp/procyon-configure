package configure

type DataSourceProperties struct {
	URL      string `yaml:"url" json:"url"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

func newDataSourceProperties() *DataSourceProperties {
	return &DataSourceProperties{}
}

func (properties *DataSourceProperties) GetConfigurationPrefix() string {
	return "procyon.datasource"
}
