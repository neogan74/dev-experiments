package designtaskmanager

import "testing"

func TestTaskManager(t *testing.T) {
	t.Run("constructorExecOrder", func(t *testing.T) {
		tasks := [][]int{
			{1, 101, 5},
			{2, 102, 7},
			{3, 103, 7},
		}

		tm := Constructor(tasks)
		tm.h.top()

		if user := tm.ExecTop(); user != 3 {
			t.Fatalf("expected user 3, got %d", user)
		}
		if user := tm.ExecTop(); user != 2 {
			t.Fatalf("expected user 2, got %d", user)
		}
		if user := tm.ExecTop(); user != 1 {
			t.Fatalf("expected user 1, got %d", user)
		}
		if user := tm.ExecTop(); user != -1 {
			t.Fatalf("expected user -1 when empty, got %d", user)
		}
	})

	t.Run("addRemove", func(t *testing.T) {
		tm := Constructor(nil)

		tm.Add(10, 1001, 4)
		tm.Add(11, 1002, 6)
		tm.Rmv(1002)

		if user := tm.ExecTop(); user != 10 {
			t.Fatalf("expected remaining user 10, got %d", user)
		}
		if user := tm.ExecTop(); user != -1 {
			t.Fatalf("expected empty manager, got user %d", user)
		}
	})

	t.Run("editUpdatesPriority", func(t *testing.T) {
		tm := Constructor([][]int{
			{21, 2001, 1},
			{22, 2002, 2},
		})

		tm.Edit(2001, 5)
		tm.Edit(2002, 3)

		if user := tm.ExecTop(); user != 21 {
			t.Fatalf("expected edited user 21 first, got %d", user)
		}
		if user := tm.ExecTop(); user != 22 {
			t.Fatalf("expected remaining user 22, got %d", user)
		}
		if user := tm.ExecTop(); user != -1 {
			t.Fatalf("expected empty manager after edits, got %d", user)
		}
	})
}
