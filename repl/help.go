package repl

import "fmt"

var commandHelpMap = map[string]string{
	CmdAddWarehouse:  "add_warehouse [type] - Adds a warehouse. Type: 'n' (normal) or 'c' (crate-enabled). Defaults to 'n'.",
	CmdShowWarehouse: "show_warehouse - Lists all warehouses with their type.",
	CmdAddRobot:      "add_robot W<id> [type] - Adds robot to warehouse W<id>. Type: 'N' (normal) or 'D' (diagonal).",
	CmdShowRobots:    "show_robots W<id> - Lists all robots in warehouse W<id>.",
	CmdMoveRobot:     "move_robot W<id> R<id> <cmds> - Commands robot R<id>. Cmds: N, S, E, W, NE, NW, SE, SW, G, D.",
	CmdAddCrate:      "add_crate W<id> x y - Adds a crate to (x, y) in a crate-enabled warehouse.",
	CmdShowCrates:    "show_crates W<id> - Lists all crate locations in the warehouse.",
	CmdShowGrid:      "show_grid W<id> - Displays robot and crate grid layout.",
	CmdShowTasks:     "show_tasks W<id> R<id> - Shows tasks for robot R<id>.",
	CmdCancelTask:    "cancel_task W<id> R<id> TASKID - Cancels task for robot.",
	CmdHelp:           "help [command] - Displays help for all or a specific command.",
	CmdExit:          "exit - Exits the program.",
}

func handleHelpCommands(parts []string) bool {
	if len(parts) == 1 {
		fmt.Println("\n Help Menu")
		for _, helpText := range commandHelpMap {
			fmt.Println("  " + helpText)
		}
		return true
	}

	cmd := parts[1]
	if help, ok := commandHelpMap[cmd]; ok {
		fmt.Println(help)
	} else {
		fmt.Printf("No help available for command: %s\n", cmd)
	}
	return true
}
