package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost        string `mapstructure:"DB_HOST"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	DBPort        string `mapstructure:"DB_PORT"`
	DATABASE      string `mapstructure:"DATABASE"`
	JWT           string `mapstructure:"SECRET_KEY"`
	AUTHTOKEN     string `mapstructure:"TWILIO_AUTH_TOKEN"`
	ACCOUNTSID    string `mapstructure:"TWILIO_ACCOUNT_SID"`
	SERVICESID    string `mapstructure:"TWILIO_SERVICES_ID"`
	USERNAME      string `mapstructure:"USER_NAME"`
	ADMINPASSWORD string `mapstructure:"ADMIN_PASSWORD"`
}

var envs = []string{
	"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_PORT", "TWILIO_AUTH_TOKEN", "TWILIO_ACCOUNT_SID", "TWILIO_SERVICES_ID","USER_NAME","ADMIN_PASSWORD",
}
var config Config

func LoadConfig() (Config, error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env") // set the file name and path
	viper.SetConfigType("env")  // set the file type

	err := viper.ReadInConfig()
	if err != nil { // handle errors while reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}
	return config, nil

}
func GetAdminDetails()(string,string){
	return config.USERNAME,config.ADMINPASSWORD
}
func GetConfig() Config {
	return config
}

// to get the secret code for JWT
func GetJWTConfig() string {
	return config.JWT
}
