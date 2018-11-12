package config

type Config struct {
	StorageSharding int `conf:"storage_sharding"` //分片数
}

func NewConfig() *Config {
	return &Config{
		StorageSharding: 127,
	}
}
