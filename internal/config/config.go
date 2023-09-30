package config

import (
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ExportPath    string `yaml:"export_path"`
	ImportPath    string `yaml:"import_path"`
	FileExtension string `yaml:"file_extension"`
	FFMPEGPath    string `yaml:"ffmpeg_path"`
	VideoCodec    string `yaml:"video_codec"`
}

func Setup(logger *zap.Logger) *Config {
	yamlCfg, err := loadYamlConfig()
	if err != nil {
		logger.Error("yaml config not found", zap.Error(err))
	}

	return yamlCfg
}

func loadYamlConfig() (*Config, error) {
	data, err := os.ReadFile("../configs/config.yaml")
	if err != nil {
		return nil, err
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
