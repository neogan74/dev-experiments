package todo

import (
	"testing"
)

func TestAdd(t *testing.T) {
	l := List{}

	taskName := "New Task !"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expeceted %q, got %q instead", taskName, l[0].Task)
	}
}

func TestCompleate(t *testing.T) {
	l := List{}

	taskName := "Create one new thing"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expeceted %q, got %q instead", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not be completed")
	}
	l.Complete(1)

	if !l[0].Done {
		t.Errorf("Task should be completed")
	}
}
