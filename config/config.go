package config

type Config struct {
	PostgresConfig `mapstructure:"postgres"`
	RedisConfig `mapstructure:"redis"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	PPassword string `mapstructure:"password"`
	DBname   string `mapstructure:"dbname"`
	Port     int    `mapstructure:"port"`
}

type RedisConfig struct {
	Addr string `mapstructure:"addr"`
	RPassword string `mapstructure:"password"`
	DB int `mapstructure:"DB"`
}
