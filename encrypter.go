package pass_generator

import (
	"code.google.com/p/go.crypto/pbkdf2"
	"code.google.com/p/go.crypto/twofish"
	"crypto/rand"
	"crypto/sha256"
)

const salt_size int = 16
const key_len int = 32

var Iter int = 1650

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}

func generate_salt() []byte {
	salt := make([]byte, salt_size)
	rand.Read(salt)
	return salt
}

func Encrypt(new_pass string, master []byte) []byte {
	c, err := twofish.NewCipher(master)
	check_err(err)
	encrypted_pass := make([]byte, 16)
	c.Encrypt(encrypted_pass, []byte(new_pass))

	return encrypted_pass
}

func Decrypt(pass, master []byte) []byte {
	c, err := twofish.NewCipher(master)
	check_err(err)

	decrypted_pass := make([]byte, 16)
	c.Decrypt(decrypted_pass, pass)

	return decrypted_pass
}

func Hash_password(pass string, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = generate_salt()
	}
	password := pbkdf2.Key([]byte(pass), salt, Iter, key_len, sha256.New)
	return password, salt
}
