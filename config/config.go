package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server        Server        `yaml:"server"`
	SqlDatabase   SqlDatabase   `yaml:"sqlDatabase"`
	NoSqlDatabase NoSqlDatabase `yaml:"noSqlDatabase"`
	Redis         Redis         `yaml:"redis"`
	App           App           `yaml:"app"`
}

type Server struct {
	Port int `yaml:"port"`
}

type SqlDatabase struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	Name            string `yaml:"name"`
	Sslmode         string `yaml:"sslmode"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
}

type NoSqlDatabase struct {
	Uri string `yaml:"uri"`
}

type Redis struct {
	Addr          string `yaml:"addr"`
	DB            int    `yaml:"db"`
	Password      string `yaml:"password"`
	LogTTLSeconds int    `yaml:"logTTLSeconds"`
}

type App struct {
	Secret        string `yaml:"secret"`
	TokenTtlHours int    `yaml:"tokenTtlHours"`
}

func MustLoad() (*Config, error) {
	f, err := os.Open("config/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("open config: %w", err)
	}
	defer f.Close()

	var cfg Config
	d := yaml.NewDecoder(f)
	if err := d.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("decode config: %w", err)
	}
	return &cfg, nil
}
