package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"



	"github.com/niranjanorkat/robot-challenge/librobot"

)

// ─── Command Constants ────────────────────────────────────────
const (
	CmdAddWarehouse  = "add_warehouse"
	CmdShowWarehouse = "show_warehouse"
	CmdAddRobot      = "add_robot"
	CmdShowRobots    = "show_robots"
	CmdMoveRobot     = "move_robot"
	CmdAddDRobot = "add_drobot"
	CmdShowGrid = "show_grid"
	CmdAddCrate     = "add_crate"
	CmdAddCWarehouse = "add_cwarehouse"
	CmdShowCrates       = "show_crates"
	CmdShowTasks  = "show_tasks"
	CmdCancelTask = "cancel_task"
	

	CmdExit          = "exit"
)

// ─── Message Constants ────────────────────────────────────────
const (
	MsgWelcome            = "Robot Simulator CLI"
	MsgPrompt             = "rcli> "
	MsgExit               = "Exiting..."

	// General Command Info
	MsgCommandList = "Commands: "

	// Warehouse
	MsgWarehouseAdded     = "Warehouse added successfully."
	MsgNoWarehouses       = "No warehouses available."
	MsgCurrentWarehouses  = "Current Warehouses:"
	MsgInvalidWarehouseID = "Invalid warehouse ID."
	
	// Crate Warehouse
	MsgCrateAdded       = "Crate added successfully."
	MsgInvalidAddCrate  = "Usage: add_crate W<id> x y"
	MsgCrateError       = "Error adding crate:"
	MsgCWarehouseAdded  = "Crate-enabled warehouse added successfully."

	// Crates
	MsgInvalidShowCrate = "Usage: show_crates W<id>"
	MsgCrateListHeader  = "Crates in warehouse:"
	MsgNoCrates         = "No crates in this warehouse."
	MsgNotCrateWarehouse = "This warehouse does not support crates."

	// Robot
	MsgRobotAdded         = "Robot added to warehouse."
	MsgRobotListHeader    = "Robots in warehouse:"
	MsgInvalidRobotID     = "Invalid robot ID."
	MsgInvalidAddRobotArgs = "Usage: add_robot W<id> x y"
	MsgInvalidShowArgs    = "Usage: show_robots W<id>"
	MsgMoveSuccess        = "Move command sent successfully."
	MsgMoveError          = "Error sending move command:"

	// Grid
	MsgInvalidShowGridArgs = "Usage: show_grid W<id>"

	// Misc
	MsgUnknownCommand = "Unknown command. Try: " 
)
// getAllCommands returns all supported commands as a single string.
func getAllCommands() string {
	return strings.Join([]string{
		CmdAddWarehouse,
		CmdShowWarehouse,
		CmdAddRobot,
		CmdShowRobots,
		CmdMoveRobot,
		CmdExit,
	}, ", ")
}

var warehouses []librobot.Warehouse

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(MsgWelcome)
	fmt.Println(MsgCommandList)

	for {
		fmt.Print(MsgPrompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		parts := strings.Fields(strings.TrimSpace(input))
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		// ─── Warehouse Commands ──────────────────────────────
		case CmdAddWarehouse:
			warehouses = append(warehouses, librobot.NewWarehouse())
			fmt.Println(MsgWarehouseAdded)

		case CmdShowWarehouse:
			if len(warehouses) == 0 {
				fmt.Println(MsgNoWarehouses)
				continue
			}
			fmt.Println(MsgCurrentWarehouses)
			for i := range warehouses {
				fmt.Printf("- W%d\n", i+1)
			}
		case CmdAddCWarehouse:
			warehouses = append(warehouses, librobot.NewCrateWarehouse())
			fmt.Println(MsgCWarehouseAdded)

		case CmdAddCrate:
			if len(parts) != 4 {
				fmt.Println(MsgInvalidAddCrate)
				continue
			}
			wID, err1 := strconv.Atoi(strings.TrimPrefix(parts[1], "W"))
			x, err2 := strconv.Atoi(parts[2])
			y, err3 := strconv.Atoi(parts[3])
			if err1 != nil || err2 != nil || err3 != nil || wID <= 0 || wID > len(warehouses) {
				fmt.Println(MsgInvalidAddCrate)
				continue
			}
			if cw, ok := warehouses[wID-1].(librobot.CrateWarehouse); ok {
				err := cw.AddCrate(uint(x), uint(y))
				if err != nil {
					fmt.Println(MsgCrateError, err)
				} else {
					fmt.Println(MsgCrateAdded)
				}
			} else {
				fmt.Println("This warehouse does not support crates.")
			}
		// Crate
		case CmdShowCrates:
			if len(parts) != 2 {
				fmt.Println(MsgInvalidShowCrate)
				continue
			}
			wID, err := strconv.Atoi(strings.TrimPrefix(parts[1], "W"))
			if err != nil || wID <= 0 || wID > len(warehouses) {
				fmt.Println(MsgInvalidWarehouseID)
				continue
			}
			warehouse := warehouses[wID-1]
			cw, ok := warehouse.(librobot.CrateWarehouse)
			if !ok {
				fmt.Println(MsgNotCrateWarehouse)
				continue
			}
			fmt.Println(MsgCrateListHeader)
			found := false
			for x := 0; x < 10; x++ {
				for y := 0; y < 10; y++ {
					if cw.HasCrate(uint(x), uint(y)) {
						fmt.Printf("- Crate at (x=%d, y=%d)\n", x, y)
						found = true
					}
				}
			}
			if !found {
				fmt.Println(MsgNoCrates)
			}

		// ─── Robot Commands ─────────────────────────────────
		case CmdAddRobot:
			if len(parts) < 2 {
				fmt.Println(MsgInvalidAddRobotArgs)
				continue
			}
			wID, err := strconv.Atoi(strings.TrimPrefix(parts[1], "W"))
			if err != nil || wID <= 0 || wID > len(warehouses) {
				fmt.Println(MsgInvalidWarehouseID)
				continue
			}
			var robotType string
			if len(parts) >= 3 {
				robotType = strings.ToUpper(parts[2])
			}

			_, err = warehouses[wID-1].AddRobot(robotType)
			if err != nil {
				fmt.Println("Error adding robot:", err)
			} else {
				fmt.Println(MsgRobotAdded)
			}

		case CmdShowRobots:
			if len(parts) != 2 {
				fmt.Println(MsgInvalidShowArgs)
				continue
			}
			wID, err := strconv.Atoi(strings.TrimPrefix(parts[1], "W"))
			if err != nil || wID <= 0 || wID > len(warehouses) {
				fmt.Println(MsgInvalidWarehouseID)
				continue
			}
			robots := warehouses[wID-1].Robots()
			if len(robots) == 0 {
				fmt.Println("No robots in this warehouse.")
				continue
			}
			fmt.Println(MsgRobotListHeader)
			for i, r := range robots {
				state := r.CurrentState()
				fmt.Printf("R%d → (x=%d, y=%d)\n", i+1, state.X, state.Y)
			}

		case CmdMoveRobot:
			if len(parts) < 4 {
				fmt.Println("Usage: move_robot W<id> R<id> <command_sequence>")
				continue
			}
			wID, err1 := strconv.Atoi(strings.TrimPrefix(parts[1], "W"))
			rID, err2 := strconv.Atoi(strings.TrimPrefix(parts[2], "R"))
			command := strings.Join(parts[3:], "")
			if err1 != nil || wID <= 0 || wID > len(warehouses) {
				fmt.Println(MsgInvalidWarehouseID)
				continue
			}
			warehouse := warehouses[wID-1]
			if err2 != nil || rID <= 0 || rID > len(warehouse.Robots()) {
				fmt.Println(MsgInvalidRobotID)
				continue
			}
			err := warehouse.SendCommand(rID-1, command)
			if err != nil {
				fmt.Println(MsgMoveError, err)
			} else {
				fmt.Println(MsgMoveSuccess)
			}
		
		

		case CmdShowGrid:
			if len(parts) != 2 {
				fmt.Println(MsgInvalidShowGridArgs)
				continue
			}
			wID, err := strconv.Atoi(strings.TrimPrefix(parts[1], "W"))
			if err != nil || wID <= 0 || wID > len(warehouses) {
				fmt.Println(MsgInvalidWarehouseID)
				continue
			}
			warehouse := warehouses[wID-1]

			// ROBOT GRID
			fmt.Println("Robot Grid:")
			robotGrid := [10][10]string{}
			for y := 0; y < 10; y++ {
				for x := 0; x < 10; x++ {
					robotGrid[y][x] = "."
				}
			}
			for i, r := range warehouse.Robots() {
				s := r.CurrentState()
				if s.Y < 10 && s.X < 10 {
					prefix := "R"
					robotGrid[s.Y][s.X] = fmt.Sprintf("%s%d", prefix, i+1)
				}
			}
			for y := 0; y < 10; y++ {
				for x := 0; x < 10; x++ {
					fmt.Printf("%-3s", robotGrid[y][x])
				}
				fmt.Println()
			}

			// CRATE GRID
			if cw, ok := warehouse.(librobot.CrateWarehouse); ok {
				fmt.Println("\nCrate Grid:")
				crateGrid := [10][10]string{}
				for y := 0; y < 10; y++ {
					for x := 0; x < 10; x++ {
						if cw.HasCrate(uint(x), uint(y)) {
							crateGrid[y][x] = "C"
						} else {
							crateGrid[y][x] = "."
						}
					}
				}
				for y := 0; y < 10; y++ {
					for x := 0; x < 10; x++ {
						fmt.Printf("%-3s", crateGrid[y][x])
					}
					fmt.Println()
				}
			}
		
		case CmdShowTasks:
			if len(parts) != 3 {
				fmt.Println("Usage: show_tasks W<id> R<id>")
				continue
			}
			wID, err1 := strconv.Atoi(strings.TrimPrefix(parts[1], "W"))
			rID, err2 := strconv.Atoi(strings.TrimPrefix(parts[2], "R"))
			if err1 != nil || wID <= 0 || wID > len(warehouses) {
				fmt.Println(MsgInvalidWarehouseID)
				continue
			}
			warehouse := warehouses[wID-1]
			if err2 != nil || rID <= 0 || rID > len(warehouse.Robots()) {
				fmt.Println(MsgInvalidRobotID)
				continue
			}
			robot := warehouse.Robots()[rID-1]
			tasks := robot.GetActiveTasks()
			if len(tasks) == 0 {
				fmt.Println("No active tasks.")
				continue
			}
			fmt.Println("Active Tasks:")
			for _, t := range tasks {
				fmt.Printf("- ID: %s | Status: %s | Commands: %s\n", t.ID, t.Status, t.RawCommand)
			}
		case CmdCancelTask:
			if len(parts) != 4 {
				fmt.Println("Usage: cancel_task W<id> R<id> TASKID")
				continue
			}
			wID, err1 := strconv.Atoi(strings.TrimPrefix(parts[1], "W"))
			rID, err2 := strconv.Atoi(strings.TrimPrefix(parts[2], "R"))
			taskID := parts[3]
			if err1 != nil || wID <= 0 || wID > len(warehouses) {
				fmt.Println(MsgInvalidWarehouseID)
				continue
			}
			warehouse := warehouses[wID-1]
			if err2 != nil || rID <= 0 || rID > len(warehouse.Robots()) {
				fmt.Println(MsgInvalidRobotID)
				continue
			}
			robot := warehouse.Robots()[rID-1]
			err := robot.CancelTask(taskID)
			if err != nil {
				fmt.Println("Failed to cancel task:", err)
			} else {
				fmt.Println("Task cancelled successfully.")
			}
		
		// ─── Exit ───────────────────────────────────────────
		case CmdExit:
			fmt.Println(MsgExit)
			return

		default:
			fmt.Println(MsgUnknownCommand)
		

		}
	}
}