package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Storage struct {
	Type string `yaml:"type"`
}

type NodeChecker struct {
	NodeAddress   string `yaml:"node_address"`
	IntervalCheck int    `yaml:"interval_check"`
}

type HTTPServer struct {
	URL string `yaml:"url"`
}

type ParserConfig struct {
	HTTPServer  `yaml:"http_server"`
	Storage     `yaml:"storage"`
	NodeChecker `yaml:"node_checker"`
}

func LoadParserConfig(path string) (*ParserConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config struct {
		Parser ParserConfig `yaml:"parser"`
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config.Parser, nil
}
