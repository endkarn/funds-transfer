package config


import (
	"bytes"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Username string `yamal: "username"`
	Password string `yamal: "password"`
	Database string `yamal: "database"`
	Host     string `yamal: "host"`
	Port     string `yamal: "port"`
}

func (config Config) GetURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
}

func SetupConfig() (Config, error) {
	var config Config
	environment := "development"
	configFile, err := ioutil.ReadFile(fmt.Sprintf(`config/%s.yml`, environment))
	if err != nil {
		fmt.Printf("Cannot read config file %s", err)
		return config, err
	}
	err = yaml.NewDecoder(bytes.NewReader(configFile)).Decode(&config)
	if err != nil {
		fmt.Printf("Cannot decode config file %s", err)
		return config, err
	}
	return config, nil
}