package benchmarks

import (
	"testing"
)

var msg = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx111111111111111111111111112222222222222222222222"

func BenchmarkWithString(b *testing.B) {
	b.Run("Zap", func(b *testing.B) {
		logger := newZapLogger()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(msg)
			}
		})
	})

	b.Run("YmLog-1024", func(b *testing.B) {
		logger := newYmLog1024Logger()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.InfoString(msg)
			}
		})
	})

	b.Run("YmLog-2048", func(b *testing.B) {
		logger := newYmLog2048Logger()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.InfoString(msg)
			}
		})
	})

}
