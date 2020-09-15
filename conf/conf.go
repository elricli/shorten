package conf

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config is all conf struct.
type Config struct {
	Redis       Redis       `yaml:"redis"`
	BloomFilter BloomFilter `yaml:"bloom_filter"`
}

// Redis config.
type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// BloomFilter config.
type BloomFilter struct {
	ExpectedInsertions uint32  `yaml:"expected_insertions"`
	FPP                float64 `yaml:"fpp"`
}

// Init yaml content.
func Init(path string) *Config {
	if path == "" {
		log.Fatalln("config path is empty")
	}
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(path, ":", err)
	}
	c := &Config{}
	if err := yaml.NewDecoder(f).Decode(c); err != nil {
		log.Fatalln(err)
	}
	return c
}
