package tests

import (
	"testing"
	"time"

	"github.com/niranjanorkat/robot-challenge/librobot"
)

func TestAddCrateTwice(t *testing.T) {
	w := librobot.NewCrateWarehouse()
	err := w.AddCrate(0, 0)
	if err != nil {
		t.Fatalf("unexpected error adding crate: %v", err)
	}

	err = w.AddCrate(0, 0)
	if err == nil {
		t.Fatal("expected error adding crate to same location, got none")
	}
}

func TestCratePickupAndDrop(t *testing.T) {
	w := librobot.NewCrateWarehouse()
	err := w.AddCrate(0, 1)
	if err != nil {
		t.Fatalf("error setting up crate: %v", err)
	}

	r, err := w.AddRobot("N")
	if err != nil {
		t.Fatalf("unexpected error adding robot: %v", err)
	}

	r.EnqueueTask("SGD") // Move South, Grab, Drop
	time.Sleep(4 * time.Second)

	state := r.CurrentState()
	cw := w.(librobot.CrateWarehouse)

	if !cw.HasCrate(state.X, state.Y) {
		t.Fatal("expected crate to be dropped at robot's current position")
	}

	if state.IsCarryingCrate {
		t.Fatal("expected robot to no longer be carrying crate")
	}
}

func TestCratePickupFailsIfOccupied(t *testing.T) {
	w := librobot.NewCrateWarehouse()
	_ = w.AddCrate(0, 1)
	_ = w.AddCrate(0, 2)

	r, err := w.AddRobot("N")
	if err != nil {
		t.Fatalf("unexpected error adding robot: %v", err)
	}

	r.EnqueueTask("SGSD") // Move South, Grab, South, Drop where crate exists
	time.Sleep(5 * time.Second)

	state := r.CurrentState()

	if state.Y != 2 {
		t.Fatalf("expected robot at Y=2, got %d", state.Y)
	}
	// If drop failed, robot should still be carrying crate
	if !state.IsCarryingCrate {
		t.Fatal("expected robot to still be carrying crate")
	}
}
