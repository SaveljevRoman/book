package congig

import "github.com/spf13/viper"

type Config struct {
	Port  string `manufacture:"PORT"`
	DBUrl string `manufacture:"DB_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/common/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	if err = viper.Unmarshal(&c); err != nil {
		return
	}

	return
}
