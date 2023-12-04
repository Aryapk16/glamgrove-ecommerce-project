package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBPort     string `mapstructure:"DB_PORT"`

	JWT_SECRET_KEY string `mapstructure:"JWT_SECRET_KEY"`
	// AccountSid       string `mapstructure:"ACCOUNT_SID"`
	// AuthToken        string `mapstructure:"AUTH_TOKEN"`
	// VerifyServiceSid string `mapstructure:"VERIFY_SERVICE_ID"`
}

// to hold all names of env variables
var envsNames = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PASSWORD", "DB_PORT", "JWT_SECRET_KEY",
}
var config Config

func LoadConfig() (config Config, err error) {

	viper.AddConfigPath("./")   // add the config path
	viper.SetConfigFile(".env") // set up the file name to viper
	err = viper.ReadInConfig()  // read the env file

	// range through through the envNames and take each envName and bind that env variable to viper
	for _, env := range envsNames {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	// Validate the Config struct using the validator package
	validate := validator.New()
	if err := validate.Struct(&config); err != nil {
		return config, err
	}

		// Unmarshal the configuration into the Config struct
		if err := viper.Unmarshal(&config); err != nil {
			return config, err
		}

	//successfully loaded the env values into struct config
	return config, nil
}

func GetJWTConfig() string {
	//config := &Config{}
	return config.JWT_SECRET_KEY
}

// func GetTwilioconfig() (string, string, string) {
// 	return config.AccountSid, config.AuthToken, config.VerifyServiceSid

// }
