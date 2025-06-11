package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"drone-deploy-technical-assessment/librobot"
)

// ─── Command Constants ────────────────────────────────────────
const (
	CmdAddWarehouse  = "add_warehouse"
	CmdShowWarehouse = "show_warehouse"
	CmdAddRobot      = "add_robot"
	CmdShowRobots    = "show_robots"
	CmdMoveRobot     = "move_robot"
	CmdShowGrid = "show_grid"

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

		// ─── Robot Commands ─────────────────────────────────
		case CmdAddRobot:
			if len(parts) < 2 {
				fmt.Println(MsgInvalidAddRobotArgs)
				continue
			}
			wID, err1 := strconv.Atoi(strings.TrimPrefix(parts[1], "W"))
			if err1 != nil || wID <= 0 || wID > len(warehouses) {
				fmt.Println(MsgInvalidWarehouseID)
				continue
			}

			_, err := warehouses[wID-1].AddRobot()
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
			grid := [10][10]string{}
			for y := 0; y < 10; y++ {
				for x := 0; x < 10; x++ {
					grid[y][x] = "."
				}
			}
			for i, r := range warehouses[wID-1].Robots() {
				s := r.CurrentState()
				if s.Y < 10 && s.X < 10 {
					grid[s.Y][s.X] = fmt.Sprintf("R%d", i+1)
				}
			}
			for y := 0; y < 10; y++ {
				for x := 0; x < 10; x++ {
					fmt.Printf("%-3s", grid[y][x])
				}
				fmt.Println()
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

