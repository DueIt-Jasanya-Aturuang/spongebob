package config

import (
	"testing"
)

func BenchmarkEnvVariable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EnvInit()
	}
}
