package config

var (
	ServerPort = ":8080"
	SecretKey  = []byte("your-256-bit-secret")
	AllowedIPs = map[string]bool{"127.0.0.1": true}
)
