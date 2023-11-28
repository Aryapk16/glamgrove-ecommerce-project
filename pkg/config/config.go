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
}

// to hold all names of env variables
var envsNames = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PASSWORD", "DB_PORT", "JWT_SECRET_KEY",
}

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

	// then unmarshel the viper into config variable
	if err := viper.Unmarshal(&config); err != nil {
		return config, err // error when unmarsheling the viper to env
	}

	// atlast validate the config file using validator pakage
	// create instance and direct validte
	if err := validator.New().Struct(config); err != nil {
		return config, err // error when validating struct
	}

	//successfully loaded the env values into struct config
	return config, nil
}

func GetJWTCofig() string {
	config := &Config{}
	return config.JWT_SECRET_KEY
}
