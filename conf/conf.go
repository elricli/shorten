package conf

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Conf is all conf struct.
type Conf struct {
	Redis RedisConf `yaml:"redis"`
}

// RedisConf .
type RedisConf struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// Init yaml content.
func Init(path string) *Conf {
	if path == "" {
		log.Fatalln("config path is empty")
	}
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(path, ":", err)
	}
	c := &Conf{}
	if err := yaml.NewDecoder(f).Decode(c); err != nil {
		log.Fatalln(err)
	}
	return c
}
