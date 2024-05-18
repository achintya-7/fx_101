package config

type Config struct {
	Topic       string
	MongoUrl    string
	PostgresUrl string
}

func NewConfig() *Config {
	return &Config{
		Topic:       "my-topic",
		MongoUrl:    "mongodb://localhost:27017",
		PostgresUrl: "postgres://localhost:5432",
	}
}
