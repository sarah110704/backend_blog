package config

var AllowedOrigins = []string{
	"http://localhost:6969",
	"https://localhost:6969",
	"http://127.0.0.1:6969",
	"https://127.0.0.1:6969",
	"http://localhost:5173",
	"https://localhost:5173",
	"https://tampilan-blog.vercel.app",
	"https://backendblog.up.railway.app",
}

func GetAllowedOrigins() []string {
	return AllowedOrigins
}
