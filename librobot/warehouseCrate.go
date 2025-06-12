package librobot

import (
	"fmt"
)

// ─── Interface ───────────────────────────────────────────────

type CrateWarehouse interface {
	Warehouse
	AddCrate(x uint, y uint) error
	DelCrate(x uint, y uint) error
	HasCrate(x uint, y uint) bool
}

// ─── Implementation ──────────────────────────────────────────

type crateWarehouse struct {
	*warehouse
	crates map[[2]uint]bool // crate location map
}

// ─── Constructor ─────────────────────────────────────────────

func NewCrateWarehouse() CrateWarehouse {
	return &crateWarehouse{
		warehouse: &warehouse{
			robots:   []Robot{},
			occupied: make(map[Position]bool),
		},
		crates: make(map[[2]uint]bool),
	}
}

// ─── Public Methods ──────────────────────────────────────────

func (cw *crateWarehouse) AddRobot(robotType string) (Robot, error) {
	id := fmt.Sprintf("R%d", len(cw.robots)+1)
	r, err := CreateRobot(robotType, id, cw)
	if err != nil {
		return nil, err
	}
	cw.robots = append(cw.robots, r)
	return r, nil
}

func (cw *crateWarehouse) AddCrate(x, y uint) error {
	pos := [2]uint{x, y}
	if cw.crates[pos] {
		return fmt.Errorf("crate already exists at (%d, %d)", x, y)
	}
	cw.crates[pos] = true
	return nil
}

func (cw *crateWarehouse) DelCrate(x, y uint) error {
	pos := [2]uint{x, y}
	if !cw.crates[pos] {
		return fmt.Errorf("no crate exists at (%d, %d)", x, y)
	}
	delete(cw.crates, pos)
	return nil
}

func (cw *crateWarehouse) HasCrate(x, y uint) bool {
	return cw.crates[[2]uint{x, y}]
}
