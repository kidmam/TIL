package channels_limitations_speed

import (
	"sync"
	"testing"
)

// https://tpaschalis.github.io/channels-limitations-speed/

func BenchmarkUnbufferedChannelEmptyStruct(b *testing.B) {
	ch := make(chan struct{})
	go func() {
		for {
			<-ch
		}
	}()
	for i := 0; i < b.N; i++ {
		ch <- struct{}{}

	}
}

func BenchmarkMutexLockUnlock(b *testing.B) {
	var mux sync.Mutex
	for i := 0; i < b.N; i++ {
		mux.Lock()
		mux.Unlock()
	}

}

func BenchmarkNaiveCopy(b *testing.B) {
	from := make([]byte, b.N)
	to := make([]byte, b.N)
	b.ReportAllocs()
	b.ResetTimer()
	b.SetBytes(1)
	copy(to, from)
}
