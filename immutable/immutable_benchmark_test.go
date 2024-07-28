package immutable

import (
	"testing"
)

func BenchmarkImmutableHuman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ImmutableCreateHuman()
	}
}

func BenchmarkMutableHuman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MutableCreateHuman()
	}
}
