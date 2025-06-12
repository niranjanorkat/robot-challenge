package tests


import (
	"testing"
	"github.com/niranjanorkat/robot-challenge/librobot"
)

func TestRobotConstructor(t *testing.T) {
	w := librobot.NewWarehouse()
	r := librobot.NewRobot("R1", w)
	state := r.CurrentState()
	if state.X != 0 || state.Y != 0 {
		t.Fatalf("expected robot to start at (0,0), got (%d,%d)", state.X, state.Y)
	}
}

func TestEnqueueTask_AddsTask(t *testing.T) {
	w := librobot.NewWarehouse()
	r := librobot.NewRobot("R1", w)
	taskID := r.EnqueueTask("NSE")
	tasks := r.GetActiveTasks()

	found := false
	for _, task := range tasks {
		if task.ID == taskID && task.RawCommand == "NSE" {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("expected task with ID %s and command NSE to be active", taskID)
	}
}

func TestCancelTask_RemovesTask(t *testing.T) {
	w := librobot.NewWarehouse()
	r := librobot.NewRobot("R1", w)
	taskID := r.EnqueueTask("SSSSSSSSSEEEEEEEEE")

	err := r.CancelTask(taskID)
	if err != nil {
		t.Fatalf("expected no error while cancelling task, got %v", err)
	}

	tasks := r.GetActiveTasks()
	for _, task := range tasks {
		t.Logf("  - ID: %s | Status: %s | Cmd: %s", task.ID, task.Status, task.RawCommand)
		if task.ID == taskID {
			t.Fatalf("expected task %s to be cancelled and removed from active list", taskID)
		}
	}
}

func TestCancelTask_InvalidID(t *testing.T) {
	w := librobot.NewWarehouse()
	r := librobot.NewRobot("R1", w)

	err := r.CancelTask("invalid-id")
	if err == nil {
		t.Fatalf("expected error when cancelling unknown task ID")
	}
}
