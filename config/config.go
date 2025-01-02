package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

const (
	defaultConfigFileName = ".aws-signer.json"
)

type Config struct {
	AccessKey       string `json:"accessKey"`
	SecretAccessKey string `json:"secretAccessKey"`
	AWSRegion       string `json:"awsRegion"`
	APIGateway      string `json:"apiGateway"`
}

var c *Config
var onceConfig sync.Once

func Configuration(configFileName ...string) (*Config, error) {

	onceConfig.Do(func() {

		var cfname string

		switch len(configFileName) {
		case 0:
			dirname, err := os.UserHomeDir()
			if err != nil {
				panic(fmt.Sprintf("Cannot get home dir: %s", err))
			}
			cfname = fmt.Sprintf("%s/%s", dirname, defaultConfigFileName)
		case 1:
			cfname = configFileName[0]
		default:
			panic("incorrect arguments for configuration file name")
		}

		configFile, err := os.Open(cfname)
		if err != nil {
			panic(fmt.Sprintf("failed to open config file %s: %s", cfname, err))
		}
		defer configFile.Close()

		decoder := json.NewDecoder(configFile)
		err = decoder.Decode(&c)
		if err != nil {
			panic(fmt.Sprintf("failed to decode configuration: %s", err))
		}
	})

	return c, nil

}
