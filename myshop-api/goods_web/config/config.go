package config

type goodsSrvInfo struct {
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

type GoodsWebInfo struct {
	Port           string      `mapstructure:"port"`
	GoodsServerInfo goodsSrvInfo `mapstructure:"goods_srv_info"`
	JwtInfo        jwtInfo     `mapstructure:"jwt"`
	RedisInfo      redisInfo   `mapstructure:"redis_info"`
	ConsulInfo     consulInfo  `mapstructure:"consul_info"`
}

type NacosInfo struct {
	Host      string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	NameSpaceId string `mapstructure:"name_space_id" json:"name_space_id"`
	DataId      string `mapstructure:"data_id" json:"data_id"`
	Group       string `mapstructure:"group" json:"group"`
}
