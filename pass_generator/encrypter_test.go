package pass_generator

import "testing"

var (
	pass string = "MeaB5KsaL/xkMnyy"
	salt []byte = generate_salt()
)

func BenchmarkEncrypt(b *testing.B) {
	Iter = 1000
	for i := 0; i < b.N; i++ {
		Hash_password(pass, salt)
	}
}
func BenchmarkEncrypt1(b *testing.B) {
	Iter = 1650
	for i := 0; i < b.N; i++ {
		Hash_password(pass, salt)
	}
}
func BenchmarkEncrypt2(b *testing.B) {
	Iter = 4096
	for i := 0; i < b.N; i++ {
		Hash_password(pass, salt)
	}
}
func BenchmarkEncrypt3(b *testing.B) {
	Iter = 10000
	for i := 0; i < b.N; i++ {
		Hash_password(pass, salt)
	}
}
func BenchmarkEncrypt4(b *testing.B) {
	Iter = 100000
	for i := 0; i < b.N; i++ {
		Hash_password(pass, salt)
	}
}

func BenchmarkEncrypt5(b *testing.B) {
	Iter = 20000
	for i := 0; i < b.N; i++ {
		Hash_password(pass, salt)
	}
}
