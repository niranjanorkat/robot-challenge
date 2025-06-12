package repl

import "fmt"

var commandHelpMap = map[string]string{
	CmdAddWarehouse:  	MsgUsageAddWarehouse,
	CmdShowWarehouse: 	MsgUsageShowWarehouse,
	CmdAddRobot:      MsgUsageAddRobot,
	CmdShowRobots:    MsgUsageShowRobots,
	CmdMoveRobot:     MsgUsageMoveRobot,
	CmdAddCrate:      MsgUsageAddCrate,
	CmdShowCrates:    MsgUsageShowCrates,
	CmdShowGrid:      MsgUsageShowGrid,
	CmdShowTasks:     MsgUsageShowTasks,
	CmdCancelTask:    MsgUsageCancelTask,
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
