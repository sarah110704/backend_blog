package config

import "os"

// JWTSecret digunakan untuk menandatangani JWT token
var JWTSecret = []byte(os.Getenv("JWT_SECRET"))
