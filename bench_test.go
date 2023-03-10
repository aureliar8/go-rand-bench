package benchrand

import (
	cryptorand "crypto/rand"
	"math/rand"
	"testing"
)

func BenchmarkMathRand(b *testing.B) {
	buf := make([]byte, 4096)
	for i := 0; i < b.N; i++ {
		_, err := rand.Read(buf)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCryptoRand(b *testing.B) {
	buf := make([]byte, 4096)
	for i := 0; i < b.N; i++ {
		_, err := cryptorand.Read(buf)
		if err != nil {
			b.Fatal(err)
		}
	}
}
