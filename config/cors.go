package config

var AllowedOrigins = []string{
	"http://localhost:6969",
	"http://localhost:5173",
	"https://blogger-go.netlify.app",
	"https://backendblog.up.railway.app",
}

func GetAllowedOrigins() []string {
	return AllowedOrigins
}