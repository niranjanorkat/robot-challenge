package tests

import (
	"testing"
	"github.com/niranjanorkat/robot-challenge/librobot"
)


func TestAddRobot_Success(t *testing.T) {
	warehouse := librobot.NewWarehouse()
	initialCount := len(warehouse.Robots())

	_, err := warehouse.AddRobot("N")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(warehouse.Robots()) != initialCount+1 {
		t.Fatalf("expected robot count to increase by 1, got %d", len(warehouse.Robots()))
	}

}

func TestAddRobot_InvalidType(t *testing.T) {
	warehouse := librobot.NewWarehouse()

	_, err := warehouse.AddRobot("X") // unsupported type
	if err == nil {
		t.Fatalf("expected error for invalid robot type, got nil")
	}

}

func TestRobots_ReturnsCopy(t *testing.T) {
	warehouse := librobot.NewWarehouse()
	_, _ = warehouse.AddRobot("N")

	robots := warehouse.Robots()
	robots[0] = nil

	// ensure internal slice is unaffected
	if warehouse.Robots()[0] == nil {
		t.Fatalf("expected internal robots slice to remain unchanged")
	}

}

func TestSendCommand_InvalidIndex(t *testing.T) {
	warehouse := librobot.NewWarehouse()
	err := warehouse.SendCommand(1, "N") // no robots yet
	if err == nil {
		t.Fatalf("expected error for invalid index, got nil")
	}

}

func TestSendCommand_ValidCommand(t *testing.T) {
	warehouse := librobot.NewWarehouse()
	r, err := warehouse.AddRobot("N")
	if err != nil {
		t.Fatalf("unexpected error adding robot: %v", err)
	}

	err = warehouse.SendCommand(0, "N")
	if err != nil {
		t.Fatalf("expected no error sending command, got %v", err)
	}

	if len(r.GetActiveTasks()) != 1 {
		t.Fatalf("expected 1 active task, got %d", len(r.GetActiveTasks()))
	}

}