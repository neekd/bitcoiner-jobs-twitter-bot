package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	OAuthToken      string   `yaml:"oAuthToken"`
	OAuthSecret     string   `yaml:"oAuthSecret"`
	ClientKey       string   `yaml:"clientKey"`
	ClientKeySecret string   `yaml:"clientKeySecret"`
	RSSFeedURL      string   `yaml:"rssFeedURL"`
	HashTags        []string `yaml:"hashtags"`
}

func Load() (*Config, error) {
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	var data *Config
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
