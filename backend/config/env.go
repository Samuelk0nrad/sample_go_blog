package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	MinioEndpoint        string `mapstructure:"MINIO_ENDPOINT"`
	MinioAccessKeyID     string `mapstructure:"MINIO_ACCESS_KEY_ID"`
	MinioSecretAccessKey string `mapstructure:"MINIO_SECRET_ACCESS_KEY_ID"`
	MinioUseSSL          bool   `mapstructure:"MINIO_USE_SSL"`
	MinioBucket          string `mapstructure:"MINIO_BUCKET"`
}

func NewEnv(filename string, override bool) *Env {
	env := Env{}

	viper.SetConfigFile(filename)

	if override {
		viper.AutomaticEnv()
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading environment file", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Error loading environment file", err)
	}

	return &env
}
