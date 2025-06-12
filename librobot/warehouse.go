package librobot

import (
	"fmt"
)

// Warehouse represents a simulated 2D grid environment where multiple robots can operate.
// It manages robot placement, movement, and ensures no two robots occupy the same cell.
type Warehouse interface {
	// Robots returns a slice of all robots currently in the warehouse.
	Robots() []Robot

	// AddRobot adds a new robot of the specified type to the warehouse.
	AddRobot(robotType string) (Robot, error)

	// SendCommand sends a movement command string to the robot at the given index.
	SendCommand(robotIndex int, command string) error

	isOccupied(x, y uint) bool
	updatePosition(oldX, oldY, newX, newY uint)
}

// Position represents a discrete (X, Y) coordinate in the warehouse grid.
type Position struct {
	X uint
	Y uint
}

type warehouse struct {
	robots   []Robot
	occupied map[Position]bool
}

// NewWarehouse creates and returns a new warehouse instance.
// Robots added to the warehouse are prevented from overlapping positions.
func NewWarehouse() Warehouse {
	return &warehouse{
		robots:   []Robot{},
		occupied: make(map[Position]bool),
	}
}

func (w *warehouse) AddRobot(robotType string) (Robot, error) {
	id := fmt.Sprintf("R%d", len(w.robots)+1)
	r, err := CreateRobot(robotType, id, w)
	if err != nil {
		return nil, err
	}
	w.robots = append(w.robots, r)
	return r, nil
}

func (w *warehouse) Robots() []Robot {
	return append([]Robot{}, w.robots...) // Return a copy
}

func (w *warehouse) SendCommand(robotIndex int, command string) error {
	if robotIndex < 0 || robotIndex >= len(w.robots) {
		return fmt.Errorf("invalid robot index: %d", robotIndex)
	}
	w.robots[robotIndex].EnqueueTask(command)
	return nil
}


func (w *warehouse) isOccupied(x, y uint) bool {
	return w.occupied[Position{X: x, Y: y}]
}

func (w *warehouse) updatePosition(oldX, oldY, newX, newY uint) {
	delete(w.occupied, Position{oldX, oldY})
	w.occupied[Position{newX, newY}] = true
}
