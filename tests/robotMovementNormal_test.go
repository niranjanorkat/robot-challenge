package tests

import (
	"testing"
	"time"
	"github.com/niranjanorkat/robot-challenge/librobot"

)

func TestRobotMovement_BasicDirections(t *testing.T) {
	w := librobot.NewWarehouse()
	r, _ := w.AddRobot("N")

	r.EnqueueTask("NSEW")
	time.Sleep(5 * time.Second)

	state := r.CurrentState()
	if state.X != 0 || state.Y != 0 {
		t.Fatalf("expected to return to origin (0,0), got (%d,%d)", state.X, state.Y)
	}
}

func TestRobotMovement_AbortOnOutOfBounds(t *testing.T) {
	w := librobot.NewWarehouse()
	r, _ := w.AddRobot("N")

	r.EnqueueTask("SSSSSSSSSSSEEEE") // 11 S + 4 E
	time.Sleep(12 * time.Second)


	state := r.CurrentState()
	if state.Y != 9 || state.X != 0 {
		t.Fatalf("expected Y to cap at 9 due to bounds, got (%d,%d)", state.X, state.Y)
	}
}

func TestRobotMovement_CollisionAbort(t *testing.T) {
	w := librobot.NewWarehouse()
	r1, _ := w.AddRobot("N")
	r2, _ := w.AddRobot("N")

	r1.EnqueueTask("EEEEEE")
	time.Sleep(2 * time.Second)
	r2.EnqueueTask("EEEEEEESSSSSSSSSS")

	time.Sleep(10 * time.Second)

	s1 := r1.CurrentState()
	s2 := r2.CurrentState()

	if s1.X != 6 || s1.Y != 0 {
		t.Fatalf("expected Robot1 at (6,0), got (%d,%d)", s1.X, s1.Y)
	}
	if s2.X != 5 || s2.Y != 0 {
		t.Fatalf("expected Robot2 to abort at (5,0), got (%d,%d)", s2.X, s2.Y)
	}
}

func TestRobotMovement_SimultaneousTasks(t *testing.T) {
	w := librobot.NewWarehouse()
	r1, _ := w.AddRobot("N")
	r2, _ := w.AddRobot("N")

	r1.EnqueueTask("EEEENNNN")
	r2.EnqueueTask("NNNNEEEE")

	time.Sleep(10 * time.Second)

	s1 := r1.CurrentState()
	s2 := r2.CurrentState()

	t.Logf("Robot1 final position: (%d,%d)", s1.X, s1.Y)
	t.Logf("Robot2 final position: (%d,%d)", s2.X, s2.Y)

	if s1.X == s2.X && s1.Y == s2.Y {
		t.Fatalf("robots should not overlap, both at (%d,%d)", s1.X, s1.Y)
	}
}
