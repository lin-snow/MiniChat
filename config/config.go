package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port       int    `yaml:"port"`
	ServerUrl  string `yaml:"server_url"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBHost     string `yaml:"db_host"`
	DBPort     int    `yaml:"db_port"`
	DBName     string `yaml:"db_name"`
}

var GlobalConfig *Config

func ParseConfig(filename string) *Config {

	// 获取当前可执行文件的完整路径
	// executable, err := os.Executable()
	// if err != nil {
	// 	log.Fatalf("\n\nUnable to get executable path: %+v\n\n", err)
	// }

	// // 如使用IDE调试，请改为本地路径
	// dir := filepath.Dir(executable)
	configPath := filepath.Join(".", filename)

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("\n\nNot found config file, %+v\n\n", err)
	}

	var cfg *Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("\n\nUnable to parse config file, %+v\n\n", err)
	}
	GlobalConfig = cfg
	return cfg
}
