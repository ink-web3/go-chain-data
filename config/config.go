package config

import "github.com/spf13/viper"

type Config struct {
	vp *viper.Viper
}

func NewConfig() (*Config, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("config")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Config{vp}, nil
}

func (config *Config) ReadSection(k string, v interface{}) error {
	err := config.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
