package _115_PrintFooBarAlternately

import (
	"strings"
	"sync"
	"testing"
)

func TestNewFooBar(t *testing.T) {
	got := NewFooBar(2)
	if got == nil {
		t.Fatal("NewFooBar(2) returned nil")
	}
	if got.n != 2 {
		t.Fatalf("FooBar.n = %d, want 2", got.n)
	}
	if got.fooCh == nil {
		t.Fatal("FooBar.fooCh is nil, want initialized channel")
	}
	if cap(got.fooCh) != 1 {
		t.Fatalf("cap(FooBar.fooCh) = %d, want 1", cap(got.fooCh))
	}
	if len(got.fooCh) != 1 {
		t.Fatalf("len(FooBar.fooCh) = %d, want 1", len(got.fooCh))
	}
	if got.barCh == nil {
		t.Fatal("FooBar.barCh is nil, want initialized channel")
	}
	if cap(got.barCh) != 1 {
		t.Fatalf("cap(FooBar.barCh) = %d, want 1", cap(got.barCh))
	}
	if len(got.barCh) != 0 {
		t.Fatalf("len(FooBar.barCh) = %d, want 0", len(got.barCh))
	}
}

func TestFooBarPrintsAlternately(t *testing.T) {
	const n = 3
	fb := NewFooBar(n)
	if fb == nil {
		t.Fatal("NewFooBar returned nil")
	}

	var (
		wg sync.WaitGroup
		mu sync.Mutex
		sb strings.Builder
	)

	printFoo := func() {
		mu.Lock()
		defer mu.Unlock()
		sb.WriteString("foo")
	}

	printBar := func() {
		mu.Lock()
		defer mu.Unlock()
		sb.WriteString("bar")
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		fb.Foo(printFoo)
	}()
	go func() {
		defer wg.Done()
		fb.Bar(printBar)
	}()
	wg.Wait()

	want := strings.Repeat("foobar", n)
	if got := sb.String(); got != want {
		t.Fatalf("output = %q, want %q", got, want)
	}
}
