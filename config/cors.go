package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"https://belatihan-production-989c.up.railway.app",
	"https://my-fe-theta.vercel.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
