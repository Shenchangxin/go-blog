package config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"name" json:"name"`
	User     string `mapstructure:"user" json:"user"`
	DBName   string `mapstructure:"dbname" json:"dbname"`
	Password string `mapstructure:"password" json:"password"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type EsConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      uint64 `mapstructure:"port" json:"port"`
	NameSpace string `mapstructure:"namespace" json:"namespace"`
	User      string `mapstructure:"user" json:"user"`
	Password  string `mapstructure:"password" json:"password"`
	DataId    string `mapstructure:"dataId" json:"dataId"`
	Group     string `mapstructure:"group" json:"group"`
}

type ServerConfig struct {
	Name         string       `mapstructure:"name" json:"name"`
	ConsulConfig ConsulConfig `mapstructure:"consul" json:"consul"`
	MysqlConfig  MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	EsConfig     EsConfig     `mapstructure:"es" json:"es"`
}
