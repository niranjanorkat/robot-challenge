package repl

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MsgRobotAdded       = "Robot added to warehouse."
	MsgRobotListHeader  = "Robots in warehouse:"
	MsgInvalidRobotID   = "Invalid robot ID."
	MsgInvalidAddRobot  = "Usage: add_robot W<id> [type]"
	MsgInvalidShowRobots = "Usage: show_robots W<id>"
	MsgInvalidMove      = "Usage: move_robot W<id> R<id> <command_sequence>"
	MsgMoveSuccess      = "Move command sent successfully."
	MsgMoveError        = "Error sending move command:"
)

func handleRobotCommands(parts []string) bool {
	switch parts[0] {

	case CmdAddRobot:
		if len(parts) < 2 {
			fmt.Println(MsgInvalidAddRobot)
			return true
		}
		wID, err := strconv.Atoi(strings.TrimPrefix(parts[1], "W"))
		if err != nil || !validWarehouseID(wID) {
			fmt.Println(MsgInvalidWarehouseID)
			return true
		}
		// Default to "N" if type not specified
		robotType := "N"
		if len(parts) >= 3 {
			robotType = strings.ToUpper(parts[2])
		}
		_, err = warehouses[wID-1].AddRobot(robotType)
		if err != nil {
			fmt.Println("Error adding robot:", err)
		} else {
			fmt.Println(MsgRobotAdded)
		}
		return true

	case CmdShowRobots:
		if len(parts) != 2 {
			fmt.Println(MsgInvalidShowRobots)
			return true
		}
		wID, err := strconv.Atoi(strings.TrimPrefix(parts[1], "W"))
		if err != nil || !validWarehouseID(wID) {
			fmt.Println(MsgInvalidWarehouseID)
			return true
		}
		robots := warehouses[wID-1].Robots()
		if len(robots) == 0 {
			fmt.Println("No robots in this warehouse.")
			return true
		}
		fmt.Println(MsgRobotListHeader)
		for i, r := range robots {
			state := r.CurrentState()
			fmt.Printf("R%d → (x=%d, y=%d)\n", i+1, state.X, state.Y)
		}
		return true

	case CmdMoveRobot:
		if len(parts) < 4 {
			fmt.Println(MsgInvalidMove)
			return true
		}
		wID, err1 := strconv.Atoi(strings.TrimPrefix(parts[1], "W"))
		rID, err2 := strconv.Atoi(strings.TrimPrefix(parts[2], "R"))
		command := strings.Join(parts[3:], "")
		if err1 != nil || !validWarehouseID(wID) {
			fmt.Println(MsgInvalidWarehouseID)
			return true
		}
		warehouse := warehouses[wID-1]
		if err2 != nil || rID <= 0 || rID > len(warehouse.Robots()) {
			fmt.Println(MsgInvalidRobotID)
			return true
		}
		err := warehouse.SendCommand(rID-1, command)
		if err != nil {
			fmt.Println(MsgMoveError, err)
		} else {
			fmt.Println(MsgMoveSuccess)
		}
		return true

	default:
		return false
	}
}
