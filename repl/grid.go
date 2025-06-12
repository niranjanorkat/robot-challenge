package repl

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/niranjanorkat/robot-challenge/librobot"
)

// ─── Grid Constants ────────────────────────────────────────────

const (
	MsgUsageShowGrid     = "show_grid W<id> [live] - Displays robot and crate grid layout. Add 'live' to watch."
	MsgInvalidShowGrid   = "Invalid show_grid command. Usage: " + MsgUsageShowGrid
	MsgInvalidWarehouse  = "Invalid warehouse ID."
	MsgRobotGridHeader   = "Robot Grid"
	MsgCrateGridHeader   = "Crate Grid"
	MsgGridEmpty         = "."
	MsgGridOrigin        = "O"
	MsgGridCrate         = "C"
	MsgGridRobotPrefix   = "R"
	MsgGridDiagonalPrefix = "D"
	MsgLiveSimulation    = "Live simulation mode. Press Ctrl+D to exit."
	MsgExitingLiveMode   = "Exiting live mode."
	clearScreen          = "\033[H\033[2J"
)


func handleGridCommands(parts []string) bool {
	if parts[0] != CmdShowGrid {
		return false
	}
	if len(parts) < 2 || len(parts) > 3 {
		fmt.Println(MsgInvalidShowGrid)
		return true
	}

	wIDStr := strings.TrimPrefix(parts[1], "W")
	wID, err := strconv.Atoi(wIDStr)
	if err != nil || !validWarehouseID(wID) {
		fmt.Println(MsgInvalidWarehouse)
		return true
	}

	if len(parts) == 3 && parts[2] == "live" {
		startLiveGridDisplay(warehouses[wID-1])
	} else {
		printGrid(warehouses[wID-1])
	}
	return true
}


func startLiveGridDisplay(warehouse librobot.Warehouse) {
	fmt.Println(MsgLiveSimulation)

	eofChan := make(chan struct{})
	go func() {
		buf := make([]byte, 1)
		for {
			_, err := os.Stdin.Read(buf)
			if err == io.EOF {
				close(eofChan)
				return
			}
		}
	}()

	for {
		select {
		case <-eofChan:
			fmt.Println(MsgExitingLiveMode)
			return
		default:
			fmt.Print(clearScreen)
			printGrid(warehouse)
			time.Sleep(1 * time.Second)
		}
	}
}


func printGrid(warehouse librobot.Warehouse) {
	// Robot Grid
	robotGrid := newEmptyGrid()
	for i, r := range warehouse.Robots() {
		state := r.CurrentState()
		if state.X >= librobot.WAREHOUSE_WIDTH || state.Y >= librobot.WAREHOUSE_HEIGHT {
			continue
		}

		label := fmt.Sprintf("%s%d", MsgGridRobotPrefix, i+1)
		if _, ok := r.(librobot.DiagonalRobot); ok {
			label = fmt.Sprintf("%s%d", MsgGridDiagonalPrefix, i+1)
		}
		robotGrid[state.Y][state.X] = label
	}
	robotGrid[0][0] = MsgGridOrigin

	fmt.Println(MsgRobotGridHeader)
	printGridMatrix(robotGrid)

	// Crate Grid
	if cw, ok := warehouse.(librobot.CrateWarehouse); ok {
		fmt.Printf("\n%s\n", MsgCrateGridHeader)
		crateGrid := newEmptyGrid()
		for y := 0; y < librobot.WAREHOUSE_HEIGHT; y++ {
			for x := 0; x < librobot.WAREHOUSE_WIDTH; x++ {
				if cw.HasCrate(uint(x), uint(y)) {
					crateGrid[y][x] = MsgGridCrate
				}
			}
		}
		crateGrid[0][0] = MsgGridOrigin
		printGridMatrix(crateGrid)
	}
}


func newEmptyGrid() [librobot.WAREHOUSE_HEIGHT][librobot.WAREHOUSE_WIDTH]string {
	var grid [librobot.WAREHOUSE_HEIGHT][librobot.WAREHOUSE_WIDTH]string
	for y := 0; y < librobot.WAREHOUSE_HEIGHT; y++ {
		for x := 0; x < librobot.WAREHOUSE_WIDTH; x++ {
			grid[y][x] = MsgGridEmpty
		}
	}
	return grid
}

func printGridMatrix(grid [librobot.WAREHOUSE_HEIGHT][librobot.WAREHOUSE_WIDTH]string) {
	for y := 0; y < librobot.WAREHOUSE_HEIGHT; y++ {
		for x := 0; x < librobot.WAREHOUSE_WIDTH; x++ {
			fmt.Printf("%-3s", grid[y][x])
		}
		fmt.Println()
	}
}
