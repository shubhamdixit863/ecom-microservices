package config

type Config struct {
	Port   string
	DBUrl  string
	DbName string
}

func NewConfig(port string, dbUrl string, dbName string) *Config {
	return &Config{
		Port:   port,
		DBUrl:  dbUrl,
		DbName: dbName,
	}
}
