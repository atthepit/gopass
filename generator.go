package pass_generator

import (
	"crypto/rand"
	"encoding/base64"
)

func Generate_random_password(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	pass := base64.StdEncoding.EncodeToString(b)
	return pass[0:length]
}

func Generate_strong_random_password() string {
	pass_length := 24
	return Generate_random_password(pass_length)
}
