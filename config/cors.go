package config

var AllowedOrigins = []string{
	"http://localhost:6969",
	"http://localhost:5173",
	"http://127.0.0.1:6969/",
}

func GetAllowedOrigins() []string {
	return AllowedOrigins
}