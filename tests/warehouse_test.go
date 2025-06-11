package tests

import (
	"testing"

	"drone-deploy-technical-assessment/librobot"
)

func TestAddRobot(t *testing.T) {
	warehouse := librobot.NewWarehouse()

	r1 := librobot.NewRobot(1, 2)
	initialCount := len(warehouse.Robots())

	err := warehouse.AddRobot(r1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(warehouse.Robots()) != initialCount+1 {
		t.Fatalf("expected robot count to increase by 1, got %d", len(warehouse.Robots()))
	}
}
