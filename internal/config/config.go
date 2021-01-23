package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

// Config is all config struct.
type Config struct {
	Domain   string   `yaml:"domain"`
	Redis    Redis    `yaml:"redis"`
	Database Database `yaml:"database"`
}

// Redis config.
type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password" json:"-"`
	DB       int    `yaml:"db"`
}

// Database config.
type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password" json:"-"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

// Init config content.
func Init(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err := yaml.NewDecoder(f).Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

// Dump outputs the current config information to the given Writer.
func (c *Config) Dump(w io.Writer) error {
	if _, err := fmt.Fprint(w, "config: "); err != nil {
		return err
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	return enc.Encode(c)
}

// DBConnInfo returns a DB connection string.
func (c *Config) DBConnInfo() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=5s&readTimeout=5s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8",
		c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)
}
