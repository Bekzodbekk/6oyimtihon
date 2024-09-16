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

	BudgetServiceHost string
	BudgetServicePort int
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
			Database: viper.GetString("postgres.database"),
			Username: viper.GetString("postgres.username"),
			Password: viper.GetString("postgres.password"),
		},
		Redis: Redis{
			Host: viper.GetString("redis.host"),
			Port: viper.GetInt("redis.port"),
		},

		BudgetServiceHost: viper.GetString("service.host"),
		BudgetServicePort: viper.GetInt("service.port"),
	}

	return &conf, nil
}
