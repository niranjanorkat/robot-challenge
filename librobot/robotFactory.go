package librobot

import (
	"fmt"
)

type RobotConstructor func(id string, wh Warehouse) Robot

var robotRegistry = map[string]RobotConstructor{
	"N": NewRobot,
	"D": NewDiagonalRobot,
}

func CreateRobot(robotType, id string, wh Warehouse) (Robot, error) {
	constructor, ok := robotRegistry[robotType]
	if !ok {
		return nil, fmt.Errorf("unknown robot type: %s", robotType)
	}
	return constructor(id, wh), nil
}

