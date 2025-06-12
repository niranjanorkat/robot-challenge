package repl

import "fmt"

const (
	MsgWelcome     = "Robot Simulator CLI"
	MsgPrompt      = "rcli> "
	MsgExit        = "Exiting..."
)

func HandleCommand(parts []string) bool {
	switch parts[0] {
	case CmdExit:
		fmt.Println(MsgExit)
		return true
	default:
		if handleWarehouseCommands(parts) ||
			handleCrateCommands(parts) ||
			handleRobotCommands(parts) ||
			handleGridCommands(parts) ||
			handleTaskCommands(parts) {
			return false
		}
		fmt.Println("Unknown command. Try:", AllCommands())
		return false
	}
}
