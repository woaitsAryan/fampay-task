package initializers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

type Config struct {
	PORT           string   `mapstructure:"PORT"`
	DB_HOST        string   `mapstructure:"DB_HOST"`
	DB_PORT        string   `mapstructure:"DB_PORT"`
	DB_NAME        string   `mapstructure:"DB_NAME"`
	DB_USER        string   `mapstructure:"DB_USER"`
	DB_PASSWORD    string   `mapstructure:"DB_PASSWORD"`
	REDIS_HOST     string   `mapstructure:"REDIS_HOST"`
	REDIS_PORT     string   `mapstructure:"REDIS_PORT"`
	REDIS_PASSWORD string   `mapstructure:"REDIS_PASSWORD"`
	DeveloperKeys  []string `mapstructure:"developerKey"`
}

var CONFIG Config

func LoadEnv() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&CONFIG)
	if err != nil {
		log.Fatal(err)
	}

	CONFIG.DeveloperKeys = LoadKeys()

	requiredKeys := getRequiredKeys(CONFIG)
	missingKeys := checkMissingKeys(requiredKeys, CONFIG)

	if len(missingKeys) > 0 {
		err := fmt.Errorf("following environment variables not found: %v", missingKeys)
		log.Fatal(err)
	}

}

func getRequiredKeys(config Config) []string {
	requiredKeys := []string{}
	configType := reflect.TypeOf(config)

	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		tag := field.Tag.Get("mapstructure")
		if tag != "" {
			requiredKeys = append(requiredKeys, tag)
		}
	}

	return requiredKeys
}

func LoadKeys() []string {
    var keys []string
    file, err := os.Open("keys.json")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }

    err = json.Unmarshal(data, &keys)
    if err != nil {
        log.Fatal(err)
    }

    return keys
}
func checkMissingKeys(requiredKeys []string, config Config) []string {
	missingKeys := []string{}

	configValue := reflect.ValueOf(config)
	for _, key := range requiredKeys {
		value := configValue.FieldByName(key).String()
		if value == "" {
			missingKeys = append(missingKeys, key)
		}
	}

	return missingKeys
}

func (c *Config) RotateDeveloperKey() {
	if len(c.DeveloperKeys) > 1 {
		c.DeveloperKeys = append(c.DeveloperKeys[1:], c.DeveloperKeys[0])
	}
}
