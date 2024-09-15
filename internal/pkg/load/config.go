package load

import "github.com/spf13/viper"

type Postgres struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

type Redis struct {
	Host string
	Port int
}

type Config struct {
	Postgres Postgres
	Redis    Redis

	UserServiceHost string
	UserServicePort int

	Email     string
	AccessKey string
	SecretKey string
}

func Load(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := Config{
		Postgres: Postgres{
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetInt("postgres.port"),
			Database: viper.GetString("postgres.dbname"),
			Username: viper.GetString("postgres.username"),
			Password: viper.GetString("postgres.password"),
		},
		Redis: Redis{
			Host: viper.GetString("redis.host"),
			Port: viper.GetInt("redis.port"),
		},

		UserServiceHost: viper.GetString("service.host"),
		UserServicePort: viper.GetInt("service.port"),
		Email:           viper.GetString("email"),
		AccessKey:       viper.GetString("accessKey"),
		SecretKey:       viper.GetString("auth.secretKey"),
	}

	return &conf, nil
}
