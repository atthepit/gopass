package pass_generator

import "testing"

func BenchmarkEncrypt(b *testing.B) {
	Iter = 1000
	for i := 0; i < b.N; i++ {
		Encrypt("eV1gUppsqtFcA9pY4+uzZogIJtd8B9D1")
	}
}
func BenchmarkEncrypt1(b *testing.B) {
	Iter = 1650
	for i := 0; i < b.N; i++ {
		Encrypt("eV1gUppsqtFcA9pY4+uzZogIJtd8B9D1")
	}
}
func BenchmarkEncrypt2(b *testing.B) {
	Iter = 4096
	for i := 0; i < b.N; i++ {
		Encrypt("eV1gUppsqtFcA9pY4+uzZogIJtd8B9D1")
	}
}
func BenchmarkEncrypt3(b *testing.B) {
	Iter = 10000
	for i := 0; i < b.N; i++ {
		Encrypt("eV1gUppsqtFcA9pY4+uzZogIJtd8B9D1")
	}
}
func BenchmarkEncrypt4(b *testing.B) {
	Iter = 100000
	for i := 0; i < b.N; i++ {
		Encrypt("eV1gUppsqtFcA9pY4+uzZogIJtd8B9D1")
	}
}
