package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Server struct {
	Port string `yaml:"port" env-required`
	Host string `yaml:"host" env-required`
}
type Database struct {
	Db_name   string `yaml:"db_name" env-required`
	User_name string `yaml:"user_name" env-required`
	Password  string `yaml:"password" env-required`
}
type Config struct {
	Env      string `yaml: "env" env-required="true"`
	Server   `yaml: "server" env-required="true"`
	Database `yaml: "database" env-required="true"`
}

// This is config set file
func MustLoad() *Config {
	flagValue := flag.String("env", "", "Provide the path of your enviroment")
	flag.Parse()
	configPath := *flagValue
	if configPath == "" {
		log.Fatal("Enviroment file must provided")
	}
	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("Config file not found %s", configPath)
	}
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Config file is not set properly %s", configPath)
	}
	return &cfg
}
