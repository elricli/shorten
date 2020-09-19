package config

import (
	"os"
	"strconv"
)

// Config is all config struct.
type Config struct {
	Redis       Redis
	BloomFilter BloomFilter
}

// Redis config.
type Redis struct {
	Addr     string
	Password string
	DB       int
}

// BloomFilter config.
type BloomFilter struct {
	ExpectedInsertions uint
	FPP                float64
	HashSeed           uint
}

// Init config content.
func Init() (*Config, error) {
	return &Config{
		Redis: Redis{
			Addr:     GetEnv("SHORTEN_REDIS_ADDR", "localhost:6379"),
			Password: GetEnv("SHORTEN_REDIS_PASSWORD", "password"),
			DB:       GetEnvInt("SHORTEN_REDIS_DB", 0),
		},
		BloomFilter: BloomFilter{
			ExpectedInsertions: GetEnvUint("SHORTEN_BLOOM_FILTER_EXPECTED_INSERTIONS", 1e7),
			FPP:                GetEnvFloat64("SHORTEN_BLOOM_FILTER_FPP", 0.00001),
			HashSeed:           GetEnvUint("SHORTEN_BLOOM_FILTER_HASH_SEED", 0x1),
		},
	}, nil
}

// GetEnv looks up the given key from the environment, returning its value if
// it exists, and otherwise returning the given fallback value.
func GetEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}

// GetEnvInt looks up the given key from the environment and expects an integer,
// returning the integer value if it exists, and otherwise returning the given
// fallback value.
func GetEnvInt(key string, fallback int) int {
	if valueStr, ok := os.LookupEnv(key); ok {
		if v, err := strconv.Atoi(valueStr); err == nil {
			return v
		}
	}
	return fallback
}

// GetEnvUint looks up the given key from the environment and expects an
// unsigned integer, returning the integer value if it exists,
// and otherwise returning the given fallback value.
func GetEnvUint(key string, fallback uint) uint {
	if valueStr, ok := os.LookupEnv(key); ok {
		if v, err := strconv.ParseUint(valueStr, 10, 64); err == nil {
			return uint(v)
		}
	}
	return fallback
}

// GetEnvFloat64 looks up the given key from the environment and expects a
// float64, returning the float64 value if it exists, and otherwise returning
// the given fallback value.
func GetEnvFloat64(key string, fallback float64) float64 {
	if valueStr, ok := os.LookupEnv(key); ok {
		if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
			return value
		}
	}
	return fallback
}
