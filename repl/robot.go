package repl

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MsgUsageAddRobot    = "add_robot W<id> [type] - Adds robot to warehouse W<id>. Type: 'N' (normal) or 'D' (diagonal)."
	MsgUsageShowRobots  = "show_robots W<id> - Lists all robots in warehouse W<id>."
	MsgUsageMoveRobot   = "move_robot W<id> R<id> <cmds> - Commands robot R<id>. Cmds: N, S, E, W, NE, NW, SE, SW, G, D."
	MsgRobotAdded       = "Robot added to warehouse."
	MsgRobotListHeader  = "Robots in warehouse:"
	MsgInvalidRobotID   = "Invalid robot ID."
	MsgMoveSuccess      = "Move command sent successfully."
	MsgMoveError        = "Error sending move command:"
	MsgNoRobots         = "No robots in this warehouse."
	MsgErrorAddingRobot = "Error adding robot:"
	MsgRobotPosition    = "R%d → (x=%d, y=%d)\n"
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
			fmt.Println(MsgErrorAddingRobot, err)
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
			fmt.Println(MsgNoRobots)
			return true
		}
		fmt.Println(MsgRobotListHeader)
		for i, r := range robots {
			state := r.CurrentState()
			fmt.Printf(MsgRobotPosition, i+1, state.X, state.Y)
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
