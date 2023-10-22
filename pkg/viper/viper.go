package viper

import "github.com/spf13/viper"

func LoadYAMLToStruct(filePath string, payload any) error {
	viperConfig := viper.New()
	viperConfig.SetConfigName(filePath)
	viperConfig.SetConfigType("yaml")
	viperConfig.AddConfigPath(".")

	if err := viperConfig.ReadInConfig(); err != nil {
		return err
	}

	if err := viperConfig.Unmarshal(&payload); err != nil {
		return err
	}

	return nil
}
