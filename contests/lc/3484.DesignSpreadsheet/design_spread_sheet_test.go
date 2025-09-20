// lc3484_spreadsheet_test.go
package designspreadsheet

import "testing"

func TestExampleFlow(t *testing.T) {
	ss := Constructor(3)
	if got := ss.GetValue("=5+7"); got != 12 {
		t.Fatalf("got %d want 12", got)
	}
	ss.SetCell("A1", 10)
	if got := ss.GetValue("=A1+6"); got != 16 {
		t.Fatalf("got %d want 16", got)
	}
	ss.SetCell("B2", 15)
	if got := ss.GetValue("=A1+B2"); got != 25 {
		t.Fatalf("got %d want 25", got)
	}
	ss.ResetCell("A1")
	if got := ss.GetValue("=A1+B2"); got != 15 {
		t.Fatalf("got %d want 15", got)
	}
}

func TestUnsetCellIsZero(t *testing.T) {
	ss := Constructor(2)
	if got := ss.GetValue("=X9+3"); got != 3 {
		t.Fatalf("got %d want 3", got)
	}
	ss.SetCell("Z5", 100)
	if got := ss.GetValue("=Z5+900"); got != 1000 {
		t.Fatalf("got %d want 1000", got)
	}
	ss.SetCell("C1", 7)
	ss.ResetCell("C1")
	if got := ss.GetValue("=C1+1"); got != 1 {
		t.Fatalf("got %d want 1", got)
	}
}
