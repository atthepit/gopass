package pass_generator

import (
	"code.google.com/p/go.crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha256"
)

const salt_size int = 16
const key_len int = 32

var Iter int = 1650

func generate_salt() []byte {
	salt := make([]byte, salt_size)
	rand.Read(salt)
	return salt
}

func Encrypt(pass string) ([]byte, []byte) {
	salt := generate_salt()
	password := pbkdf2.Key([]byte(pass), salt, Iter, key_len, sha256.New)
	return password, salt
}
