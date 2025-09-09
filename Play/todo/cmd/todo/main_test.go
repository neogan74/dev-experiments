package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool: %v", err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	fmt.Println("Cleaning up...")
	if err := os.Remove(binName); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot remove binary %q: %v", binName, err)
	}

	if err := os.Remove(fileName); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot remove file %q: %v", fileName, err)
	}

	os.Exit(result)
}

func TestTodoCli(t *testing.T) {
	task := "New Task from CLI"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Cannot get working directory: %v", err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("Add Task", func(t *testing.T) {
		cmd := exec.Command(cmdPath, strings.Split(task, " ")...)

		if err := cmd.Run(); err != nil {
			t.Fatalf("Command %q failed: %v", cmd.Args, err)
		}
	})
	t.Run("List Task", func(t *testing.T) {
		cmd := exec.Command(cmdPath)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Command %q failed: %v", cmd.Args, err)
		}

		expected := task + "\n"
		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead", expected, out)
		}
	})
}
