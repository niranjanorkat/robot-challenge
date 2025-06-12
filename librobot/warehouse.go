package librobot

import (
	"fmt"
)

// ─── Interface ───────────────────────────────────────────────

type Warehouse interface {
	Robots() []Robot
	AddRobot(robotType string) (Robot, error)
	SendCommand(robotIndex int, command string) error
	isOccupied(x, y uint) bool
	updatePosition(oldX, oldY, newX, newY uint)
}

// ─── Types ───────────────────────────────────────────────────

type Position struct {
	X uint
	Y uint
}

type warehouse struct {
	robots   []Robot
	occupied map[Position]bool
}

// ─── Constructor ─────────────────────────────────────────────

func NewWarehouse() Warehouse {
	return &warehouse{
		robots:   []Robot{},
		occupied: make(map[Position]bool),
	}
}

// ─── Public Methods ──────────────────────────────────────────

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

// ─── Internal Helpers ────────────────────────────────────────

func (w *warehouse) isOccupied(x, y uint) bool {
	return w.occupied[Position{X: x, Y: y}]
}

func (w *warehouse) updatePosition(oldX, oldY, newX, newY uint) {
	delete(w.occupied, Position{oldX, oldY})
	w.occupied[Position{newX, newY}] = true
}
