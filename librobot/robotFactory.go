package librobot

import (
	"fmt"
)

// RobotConstructor defines a function that creates a new Robot instance
// with the given ID and warehouse context.
type RobotConstructor func(id string, wh Warehouse) Robot

var robotRegistry = map[string]RobotConstructor{
	"N": NewRobot,
	"D": NewDiagonalRobot,
}

// CreateRobot returns a new Robot of the specified type.
// Supported types:
//   - "N": Normal robot
//   - "D": Diagonal-capable robot
//
// Returns an error if the type is unknown.
func CreateRobot(robotType, id string, wh Warehouse) (Robot, error) {
	constructor, ok := robotRegistry[robotType]
	if !ok {
		return nil, fmt.Errorf("unknown robot type: %s", robotType)
	}
	return constructor(id, wh), nil
}

