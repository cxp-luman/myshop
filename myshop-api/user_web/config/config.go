package config

type userSrvInfo struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

type consulInfo struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}

type jwtInfo struct {
	Key string `mapstructure:"key"`
}

type redisInfo struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type UserWebInfo struct {
	Port           string      `mapstructure:"port"`
	UserServerInfo userSrvInfo `mapstructure:"user_srv_info"`
	JwtInfo        jwtInfo     `mapstructure:"jwt"`
	RedisInfo      redisInfo   `mapstructure:"redis_info"`
	ConsulInfo     consulInfo  `mapstructure:"consul_info"`
}

type NacosInfo struct {
	Ipaddr      string `mapstructure:"ipaddr" json:"ipaddr"`
	Port        int    `mapstructure:"port" json:"port"`
	NameSpaceId string `mapstructure:"name_space_id" json:"name_space_id"`
	DataId      string `mapstructure:"data_id" json:"data_id"`
	Group       string `mapstructure:"group" json:"group"`
}
