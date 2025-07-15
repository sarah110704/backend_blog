package config

var AllowedOrigins = []string{
	"http://localhost:6969",
	"http://localhost:5173",
}

func GetAllowedOrigins() []string {
	return AllowedOrigins
}