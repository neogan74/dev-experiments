package main

import (
	"bytes"
	"sync"
	"testing"
)

func TestFooPrintOrder(t *testing.T) {
	for i := 0; i < 100; i++ {
		foo := Constructor()

		var buf bytes.Buffer
		var mu sync.Mutex
		var wg sync.WaitGroup
		wg.Add(3)

		go func() {
			defer wg.Done()
			foo.Third(func() {
				mu.Lock()
				buf.WriteString("third")
				mu.Unlock()
			})
		}()

		go func() {
			defer wg.Done()
			foo.First(func() {
				mu.Lock()
				buf.WriteString("first")
				mu.Unlock()
			})
		}()

		go func() {
			defer wg.Done()
			foo.Second(func() {
				mu.Lock()
				buf.WriteString("second")
				mu.Unlock()
			})
		}()

		wg.Wait()

		if got, want := buf.String(), "firstsecondthird"; got != want {
			t.Fatalf("iteration %d: got %q, want %q", i, got, want)
		}
	}
}

func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo := Constructor()

		var buf bytes.Buffer
		var mu sync.Mutex
		var wg sync.WaitGroup
		wg.Add(3)

		go func() {
			defer wg.Done()
			foo.Third(func() {
				mu.Lock()
				buf.WriteString("third")
				mu.Unlock()
			})
		}()

		go func() {
			defer wg.Done()
			foo.First(func() {
				mu.Lock()
				buf.WriteString("first")
				mu.Unlock()
			})
		}()

		go func() {
			defer wg.Done()
			foo.Second(func() {
				mu.Lock()
				buf.WriteString("second")
				mu.Unlock()
			})
		}()

		wg.Wait()

		if got, want := buf.String(), "firstsecondthird"; got != want {
			b.Fatalf("got %q, want %q", got, want)
		}
	}
}
