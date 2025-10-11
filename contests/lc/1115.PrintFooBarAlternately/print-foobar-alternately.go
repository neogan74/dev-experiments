package _115_PrintFooBarAlternately

type FooBar struct {
	n     int
	fooCh chan struct{}
	barCh chan struct{}
}

func NewFooBar(n int) *FooBar {
	f := &FooBar{
		n:     n,
		fooCh: make(chan struct{}, 1), // буфер 1, чтобы стартовать с foo
		barCh: make(chan struct{}, 1),
	}
	f.fooCh <- struct{}{} // сначала разрешаем foo
	return f
}

func (f *FooBar) Foo(printFoo func()) {
	for i := 0; i < f.n; i++ {
		<-f.fooCh
		printFoo()
		f.barCh <- struct{}{}
	}
}

func (f *FooBar) Bar(printBar func()) {
	for i := 0; i < f.n; i++ {
		<-f.barCh
		printBar()
		f.fooCh <- struct{}{}
	}
}
