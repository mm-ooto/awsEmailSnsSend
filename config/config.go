package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var ConfigData *Config

type Config struct {
	RunMode string `mapstructure:"run_mode"`
	Email   struct {
		Region      string
		AccessKeyId string `mapstructure:"access_key_id"`
		SecretKey   string `mapstructure:"secret_key"`
		Sender      string
	}
	Sms struct {
		Region      string
		AccessKeyId string `mapstructure:"access_key_id"`
		SecretKey   string `mapstructure:"secret_key"`
	}
}

func LoadConfig(runMode string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if runMode == "" {
		runMode = "dev"
	}
	path := fmt.Sprintf("./conf/%s", runMode)
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper ReadInConfig err:%s\n", err.Error())
		return
	}
	if err := viper.Unmarshal(&ConfigData); err != nil {
		log.Fatalf("viper Unmarshal err:%s\n", err.Error())
		return
	}
	bytes, _ := json.Marshal(ConfigData)
	fmt.Println("conf data:", string(bytes))
}
