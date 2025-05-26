package hw7

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// С помощью бенчмарков можно узнать не только время,
// но и много чего ещё интересного, но их раскрутка занимает больше времени,
// чем обычный time.since - time.now
func TestTimeBench(t *testing.T) {
	allowedTime := time.Millisecond * 500
	bench1res := testing.Benchmark(func(b *testing.B) {
		f1()
	})
	require.Less(t, bench1res.T, allowedTime)

	bench2res := testing.Benchmark(func(b *testing.B) {
		f2()
	})
	require.Less(t, bench2res.T, allowedTime)

	bench3res := testing.Benchmark(func(b *testing.B) {
		f3()
	})
	assert.Less(t, bench3res.T, allowedTime*2) // *2 для прохождения теста
}

// Если волнует только время
func TestTimeTimer(t *testing.T) {
	allowedTime := time.Millisecond * 500

	start := time.Now()
	f1()
	elapsed1 := time.Since(start)
	require.Less(t, elapsed1, allowedTime)

	start = time.Now()
	f2()
	elapsed2 := time.Since(start)
	require.Less(t, elapsed2, allowedTime)

	start = time.Now()
	f3()
	elapsed3 := time.Since(start)
	assert.Less(t, elapsed3, allowedTime*2) // *2 для прохождения теста
}
