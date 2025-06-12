package librobot

import (
	"fmt"
)

// CrateWarehouse extends Warehouse functionality by supporting crate placement and retrieval.
// In addition to managing robots, it tracks crate positions and allows operations on them.
type CrateWarehouse interface {
	Warehouse

	// AddCrate places a crate at the specified (x, y) position.
	AddCrate(x uint, y uint) error

	// DelCrate removes the crate from the specified position.
	DelCrate(x uint, y uint) error

	// HasCrate returns true if a crate exists at the given position.
	HasCrate(x uint, y uint) bool
}



type crateWarehouse struct {
	*warehouse
	crates map[[2]uint]bool // crate location map
}

// NewCrateWarehouse returns a new CrateWarehouse instance with empty robot and crate maps.
func NewCrateWarehouse() CrateWarehouse {
	return &crateWarehouse{
		warehouse: &warehouse{
			robots:   []Robot{},
			occupied: make(map[Position]bool),
		},
		crates: make(map[[2]uint]bool),
	}
}

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
