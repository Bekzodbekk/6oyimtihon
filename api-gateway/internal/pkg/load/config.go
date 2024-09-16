package load

import "github.com/spf13/viper"

type ServiceConfig struct {
	Host string
	Port string
}

type Config struct {
	ApiGateway       ServiceConfig
	UserService      ServiceConfig
	IncExpService    ServiceConfig
	ReportingService ServiceConfig
	BudgetService    ServiceConfig
	TokenKey         string
	CertFile         string
	KeyFile          string
}

func Load(path string) (*Config, error) {

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		ApiGateway: ServiceConfig{
			Host: viper.GetString("api-gateway.host"),
			Port: viper.GetString("api-gateway.port"),
		},
		UserService: ServiceConfig{
			Host: viper.GetString("services.user-service.host"),
			Port: viper.GetString("services.user-service.port"),
		},
		IncExpService: ServiceConfig{
			Host: viper.GetString("services.incexp-service.host"),
			Port: viper.GetString("services.incexp-service.port"),
		},
		ReportingService: ServiceConfig{
			Host: viper.GetString("services.reporting-service.host"),
			Port: viper.GetString("services.reporting-service.port"),
		},
		BudgetService: ServiceConfig{
			Host: viper.GetString("services.budget-service.host"),
			Port: viper.GetString("services.budget-service.port"),
		},

		TokenKey: viper.GetString("token.key"),
		
		CertFile: viper.GetString("file.cert"),
		KeyFile:  viper.GetString("file.key"),
	}
	return &cfg, nil
}
