package configure

type DataSourceProperties struct {
	DriverName   string `yaml:"driver-name" json:"driver-name"`
	Username     string `yaml:"username" json:"username"`
	Password     string `yaml:"password" json:"password"`
	DatabaseName string `yaml:"database-name" json:"database-name"`
}

func newDataSourceProperties() *DataSourceProperties {
	return &DataSourceProperties{}
}

func (properties *DataSourceProperties) GetPrefix() string {
	return "procyon.datasource"
}
