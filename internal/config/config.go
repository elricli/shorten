package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

// Config is all config struct.
type Config struct {
	Domain      string      `yaml:"domain"`
	Redis       Redis       `yaml:"redis"`
	BloomFilter BloomFilter `yaml:"bloom_filter"`
	Database    Database    `yaml:"database"`
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

// BloomFilter config.
type BloomFilter struct {
	ExpectedInsertions uint    `yaml:"expected_insertions"`
	FPP                float64 `yaml:"fpp"`
	HashSeed           uint    `yaml:"hash_seed"`
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

// DBConnInfo returns a PostgreSQL connection string.
func (c *Config) DBConnInfo() string {
	timeoutOption := fmt.Sprintf("-c statement_timeout=%d", 10*time.Minute/time.Millisecond)
	return fmt.Sprintf("user='%s' password='%s' host='%s' port=%s dbname='%s' sslmode=disable options='%s'",
		c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name, timeoutOption)
}
