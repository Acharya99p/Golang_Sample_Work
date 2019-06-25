package config

type Config struct {
	DB *DBConfig
	RmQ *RabbitConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Host     string
	Charset  string
}

type RabbitConfig struct {
	Dialect string
	Host string
	Username string
	Password string
	Port string
	ApiPort string
}



func GetConfig() *Config {

	return &Config{
		DB: &DBConfig{
			Dialect:  "sqlite3",
			Username: "",
			Password: "",
			Name:     "/tmp/test.db",
			Charset:  "utf8",
		},
		RmQ: &RabbitConfig{
			Dialect:  "amqp",
			Username: "guest",
			Password: "guest",
			Host:     "localhost",
			Port:     "5672",
			ApiPort:  "15672",
		},

	}
}
