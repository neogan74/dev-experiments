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

func TestDelete(t *testing.T) {
	l := List{}

	tasks := []string{
		"Task 1",
		"Task 2",
		"Task 3",
	}

	for _, task := range tasks {
		l.Add(task)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected list length %d, got %d instead", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead", tasks[2], l[1].Task)
	}
}
