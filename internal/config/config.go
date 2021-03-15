package config

type Config struct {
	HttpServer struct {
		Host string
		Port int
	}
	DataBase struct {
		Host     string
		Port     int
		User     string
		Password string
		DbName   string
	}
	Redis struct {
		Address  string
		Password string
		DB       int
	}
}
