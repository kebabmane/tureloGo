package app

import (
	"fmt"
	"log"
	"os"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

// Config stores the application-wide configurations
var Config appConfig

type appConfig struct {
	// the path to the error message file. Defaults to "config/errors.yaml"
	ErrorFile string `mapstructure:"error_file"`
	// the server port. Defaults to 8080
	ServerPort int `mapstructure:"server_port"`
	// the data source name (DSN) for connecting to the database. required.
	DSN string `mapstructure:"dsn"`
	// the signing method for JWT. Defaults to "HS256"
	JWTSigningMethod string `mapstructure:"jwt_signing_method"`
	// JWT signing key. required.
	JWTSigningKey string `mapstructure:"jwt_signing_key"`
	// JWT verification key. required.
	JWTVerificationKey string `mapstructure:"jwt_verification_key"`
}

func (config appConfig) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.DSN, validation.Required),
		validation.Field(&config.JWTSigningKey, validation.Required),
		validation.Field(&config.JWTVerificationKey, validation.Required),
	)
}

// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
// The configuration file(s) should be named as app.yaml.
// Environment variables with the prefix "PRODUCTION_" in their names are also read automatically.
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	v.SetDefault("error_file", "config/errors.yaml")

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}

	if err := v.Unmarshal(&Config); err != nil {
		return err
	}

	if os.Getenv("DATABASE_URL") == "" {
		log.Printf("$DSN not set, setting default")
	} else {
		log.Printf("Setting database via env")
		log.Printf("DATABASE_URL:", os.Getenv("DATABASE_URL"))
		v.BindEnv("dsn", "DATABASE_URL")
		log.Printf("DSN to Use: ", v.GetString("dsn"))
	}

	if os.Getenv("PORT") == "" {
		log.Printf("$PORT not set, setting from config")
	} else {
		log.Printf("Setting port via env")
		v.Set("server_port", os.Getenv("PORT"))
	}

	return Config.Validate()
}
