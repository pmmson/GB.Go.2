package task03_test

import (
	"testing"

	"github.com/pmmson/GB.Go.2/hw05/task03"
)

func BenchmarkSet10W90R(b *testing.B) {
	var set = task03.NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(0.1)
			}
		})
		b.SetParallelism(900)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(0.1)
			}
		})
	})
}

func BenchmarkSet50W50R(b *testing.B) {
	var set = task03.NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(500)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(0.1)
			}
		})
		b.SetParallelism(500)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(0.1)
			}
		})
	})
}

func BenchmarkSet90W10R(b *testing.B) {
	var set = task03.NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(900)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(0.1)
			}
		})
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(0.1)
			}
		})
	})
}

func BenchmarkSetRW10W90R(b *testing.B) {
	var set = task03.NewSetRW()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddRW(0.1)
			}
		})
		b.SetParallelism(900)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.HasRW(0.1)
			}
		})
	})
}

func BenchmarkSetRW50W50R(b *testing.B) {
	var set = task03.NewSetRW()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(500)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddRW(0.1)
			}
		})
		b.SetParallelism(500)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.HasRW(0.1)
			}
		})
	})
}

func BenchmarkSetRW90W10R(b *testing.B) {
	var set = task03.NewSetRW()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(900)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddRW(0.1)
			}
		})
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.HasRW(0.1)
			}
		})
	})
}
