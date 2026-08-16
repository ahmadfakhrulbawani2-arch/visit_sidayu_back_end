package corscfg

import "os"

var staged_url = os.Getenv("CORS_FRONT_END_STAGED_URL")

var(
	AllowedOrigins = []string{
		"http://localhost:3000",
		"http://localhost:3001",
		"http://127.0.0.1:3000",
		"http://127.0.0.1:3001",
		"http://10.73.51.153:3000",
		"http://10.73.51.153:3001",
		"http://192.168.1.9:3000",
		"http://10.46.17.153:3000",
		"https://visit-sidayu-front-end.vercel.app",
		staged_url,
	}
	AllowedMethods = []string{
		"GET",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
		"OPTIONS",
	}
)
