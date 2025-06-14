package repl

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MsgUsageShowTasks  = "show_tasks W<id> R<id> - Shows tasks for robot R<id>."
	MsgUsageCancelTask = "cancel_task W<id> R<id> TASKID - Cancels task for robot."

	MsgNoTasks        = "No active tasks."
	MsgCancelFailed   = "Failed to cancel task:"
	MsgCancelSuccess  = "Task cancelled successfully."
)

var (
	MsgInvalidShowTasks  = "Invalid show_tasks command. Usage: " + MsgUsageShowTasks
	MsgInvalidCancelTask = "Invalid cancel_task command. Usage: " + MsgUsageCancelTask
)

func handleTaskCommands(parts []string) bool {
	switch parts[0] {

	case CmdShowTasks:
		if len(parts) != 3 {
			fmt.Println(MsgInvalidShowTasks)
			return true
		}

		wIDStr := strings.TrimPrefix(parts[1], "W")
		rIDStr := strings.TrimPrefix(parts[2], "R")

		wID, err1 := strconv.Atoi(wIDStr)
		rID, err2 := strconv.Atoi(rIDStr)

		isInvalidWarehouse := err1 != nil || !validWarehouseID(wID)
		if isInvalidWarehouse {
			fmt.Println(MsgInvalidWarehouseID)
			return true
		}

		isInvalidRobot := err2 != nil || rID <= 0 || rID > len(warehouses[wID-1].Robots())
		if isInvalidRobot {
			fmt.Println(MsgInvalidRobotID)
			return true
		}

		robot := warehouses[wID-1].Robots()[rID-1]
		tasks := robot.GetActiveTasks()
		if len(tasks) == 0 {
			fmt.Println(MsgNoTasks)
			return true
		}

		fmt.Println("Active Tasks:")
		for _, t := range tasks {
			fmt.Printf("- ID: %s | Status: %s | Commands: %s\n", t.ID, t.Status, t.RawCommand)
		}
		return true

	case CmdCancelTask:
		if len(parts) != 4 {
			fmt.Println(MsgInvalidCancelTask)
			return true
		}

		wIDStr := strings.TrimPrefix(parts[1], "W")
		rIDStr := strings.TrimPrefix(parts[2], "R")
		taskID := parts[3]

		wID, err1 := strconv.Atoi(wIDStr)
		rID, err2 := strconv.Atoi(rIDStr)

		isInvalidWarehouse := err1 != nil || !validWarehouseID(wID)
		if isInvalidWarehouse {
			fmt.Println(MsgInvalidWarehouseID)
			return true
		}

		isInvalidRobot := err2 != nil || rID <= 0 || rID > len(warehouses[wID-1].Robots())
		if isInvalidRobot {
			fmt.Println(MsgInvalidRobotID)
			return true
		}

		robot := warehouses[wID-1].Robots()[rID-1]
		err := robot.CancelTask(taskID)
		if err != nil {
			fmt.Println(MsgCancelFailed, err)
		} else {
			fmt.Println(MsgCancelSuccess)
		}
		return true

	default:
		return false
	}
}