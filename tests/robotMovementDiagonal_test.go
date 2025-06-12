package tests

import (
	"testing"
	"time"

	"github.com/niranjanorkat/robot-challenge/librobot"
)

func TestDiagonalMovement(t *testing.T) {
	w := librobot.NewWarehouse()
	r, err := w.AddRobot("D")
	if err != nil {
		t.Fatalf("unexpected error adding robot: %v", err)
	}
	r.EnqueueTask("SENESESWNW")
	time.Sleep(10 * time.Second)

	pos := r.CurrentState()
	if pos.X != 1 || pos.Y != 1 {
		t.Fatalf("expected robot to return to (1,1), got (%d,%d)", pos.X, pos.Y)
	}
}