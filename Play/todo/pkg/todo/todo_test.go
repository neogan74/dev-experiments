package todo

import (
	"os"
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

func TestSaveGet(t *testing.T) {
	l1 := List{}
	l2 := List{}

	taskName := "New Task!"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task", l1[0].Task, l2[0].Task)
	}
}
