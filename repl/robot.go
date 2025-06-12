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
var (
	MsgInvalidAddRobot    = "Invalid add_robot command. Usage: " + MsgUsageAddRobot
	MsgInvalidShowRobots  = "Invalid show_robots command. Usage: " + MsgUsageShowRobots
	MsgInvalidMove        = "Invalid move_robot command. Usage: " + MsgUsageMoveRobot
)
func handleRobotCommands(parts []string) bool {
	switch parts[0] {

	case CmdAddRobot:
		if len(parts) < 2 {
			fmt.Println(MsgInvalidAddRobot)
			return true
		}

		wIDStr := strings.TrimPrefix(parts[1], "W")
		wID, err := strconv.Atoi(wIDStr)
		isInvalidWarehouse := err != nil || !validWarehouseID(wID)
		if isInvalidWarehouse {
			fmt.Println(MsgInvalidWarehouseID)
			return true
		}

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

		wIDStr := strings.TrimPrefix(parts[1], "W")
		wID, err := strconv.Atoi(wIDStr)
		isInvalidWarehouse := err != nil || !validWarehouseID(wID)
		if isInvalidWarehouse {
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

		wIDStr := strings.TrimPrefix(parts[1], "W")
		rIDStr := strings.TrimPrefix(parts[2], "R")
		command := strings.Join(parts[3:], "")

		wID, err1 := strconv.Atoi(wIDStr)
		rID, err2 := strconv.Atoi(rIDStr)

		isInvalidWarehouse := err1 != nil || !validWarehouseID(wID)
		isInvalidRobot := err2 != nil || rID <= 0 || rID > len(warehouses[wID-1].Robots())

		if isInvalidWarehouse {
			fmt.Println(MsgInvalidWarehouseID)
			return true
		}

		if isInvalidRobot {
			fmt.Println(MsgInvalidRobotID)
			return true
		}

		err := warehouses[wID-1].SendCommand(rID-1, command)
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
