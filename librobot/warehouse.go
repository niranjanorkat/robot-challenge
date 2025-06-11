package librobot

import (
	"fmt"
)
type Warehouse interface {
	Robots() []Robot
	AddRobot()(Robot, error)
	SendCommand(robotIndex int, command string) error 
	IsOccupied(x, y uint) bool
	updatePosition(oldX, oldY, newX, newY uint)
}

type Position struct {
	X uint
	Y uint
}

type warehouse struct {
	robots []Robot
	occupied map[Position]bool

}

// NewWarehouse creates a new instance of a warehouse.
func NewWarehouse() Warehouse {
	return &warehouse{
		robots: []Robot{},
		occupied: make(map[Position]bool),
	}
}

// AddRobot adds a robot to the warehouse.
func (w *warehouse) AddRobot() (Robot, error) {
	id := fmt.Sprintf("R%d", len(w.robots)+1)
	r := NewRobot(id, w)
	w.robots = append(w.robots, r)
	return r, nil
}

// Robots returns a list of robots currently in the warehouse.
func (w *warehouse) Robots() []Robot {
	return append([]Robot{}, w.robots...) // return a copy
}

// SendCommand enqueues a movement command to the robot at the given index.
func (w *warehouse) SendCommand(robotIndex int, command string) error {
	if robotIndex < 0 || robotIndex >= len(w.robots) {
		return fmt.Errorf("invalid robot index: %d", robotIndex)
	}
	_, posCh, errCh := w.robots[robotIndex].EnqueueTask(command)

	// Print live updates
	for state := range posCh {
		fmt.Printf("Robot moved to (x=%d, y=%d)\n", state.X, state.Y)
	}

	return <-errCh
}

func (w *warehouse) IsOccupied(x, y uint) bool {
	return w.occupied[Position{X: x, Y: y}]
}

func (w *warehouse) updatePosition(oldX, oldY, newX, newY uint) {
	delete(w.occupied, Position{oldX, oldY})
	w.occupied[Position{newX, newY}] = true
}