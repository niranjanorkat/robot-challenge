package repl

import "strings"

const (
	// Warehouse
	CmdAddWarehouse  = "add_warehouse"
	CmdShowWarehouse = "show_warehouse"

	// Crate
	CmdAddCrate   = "add_crate"
	CmdShowCrates = "show_crates"

	// Robot (add more as you move robot commands)
	CmdAddRobot   = "add_robot"
	CmdShowRobots = "show_robots"
	CmdMoveRobot        = "move_robot"
	// Grid
	CmdShowGrid = "show_grid"

	// Tasks
	CmdShowTasks  = "show_tasks"
	CmdCancelTask = "cancel_task"

	// Exit
	CmdExit = "exit"
)

func AllCommands() string {
	return strings.Join([]string{
		CmdAddWarehouse,
		CmdShowWarehouse,
		CmdAddCrate,
		CmdShowCrates,
		CmdAddRobot,
		CmdShowRobots,
		CmdShowGrid,
		CmdShowTasks,
		CmdCancelTask,
		CmdExit,
	}, ", ")
}
