package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	SQSCreatedOrdersQueueUrl string `mapstructure:"SQS_CREATED_ORDERS_QUEUE_URL"`
	AWSAccessKey             string `mapstructure:"AWS_ACCESS_KEY"`
	AWSSecretKey             string `mapstructure:"AWS_SECRET_KEY"`
	DatabaseURL              string `mapstructure:"DATABASE_URL"`
	WebPort                  string `mapstructure:"WEB_PORT"`
	WebHost                  string `mapstructure:"WEB_HOST"`
	Env                      string `mapstructure:"ENV"`
	DatabaseName             string `mapstructure:"DATABASE_NAME"`
}

func LoadConfig(path string) (*config, error) {
	var cfg *config

	viper.SetConfigName("ms_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Println("error on unmarshal env variables", err)
		panic(err)
	}

	return cfg, nil
}
