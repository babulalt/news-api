package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Name string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Name: "news.sqlite",
		},
	}
}
