package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig(path []string, name string, configData interface{}) (interface{}, error) {
	v := viper.NewWithOptions(viper.KeyDelimiter("__"))
	for _, value := range path {
		v.AddConfigPath(value)
	}
	v.SetConfigName(name)
	v.SetConfigType("env")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	errReadingConfig := v.ReadInConfig()
	if errReadingConfig != nil {
		fmt.Println("error data", errReadingConfig)
		return configData, errReadingConfig
	}
	errUnmarshalConfig := v.Unmarshal(&configData)
	return configData, errUnmarshalConfig
}
