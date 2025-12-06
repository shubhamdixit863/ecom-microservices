package config

type Config struct {
	Port   string
	DBUrl  string
	DbName string
}

// loading it from the env

func NewConfig() {
	// read it from enev
}
