package config

type mysqlInfo struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	DB       string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type consulInfo struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type InvSrvInfo struct {
	Name       string     `mapstructure:"name" json:"name"`
	MysqlInfo  mysqlInfo  `mapstructure:"mysql_info" json:"mysql_info"`
	ConsulInfo consulInfo `mapstructure:"consul_info" json:"consul_info"`
}

type NacosInfo struct {
	Ipaddr      string `mapstructure:"ipaddr" json:"ipaddr"`
	Port        int    `mapstructure:"port" json:"port"`
	NameSpaceId string `mapstructure:"name_space_id" json:"name_space_id"`
	DataId      string `mapstructure:"data_id" json:"data_id"`
	Group       string `mapstructure:"group" json:"group"`
}
