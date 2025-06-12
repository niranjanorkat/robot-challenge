package repl

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/niranjanorkat/robot-challenge/librobot"
)

const (
	MsgUsageShowGrid      = "show_grid W<id> [live] - Displays robot and crate grid layout. Add 'live' to watch."
	MsgRobotGridHeader    = "Robot Grid"
	MsgCrateGridHeader    = "Crate Grid"
	MsgGridEmptyCell      = "."
	MsgGridOriginCell     = "O"
	MsgGridCrateCell      = "C"
	MsgGridRobotPrefix    = "R"
	MsgGridDiagonalPrefix = "D"
)

var (
	MsgInvalidShowGrid = "Invalid show_grid command. Usage: " + MsgUsageShowGrid
)

const clearScreen = "\033[H\033[2J"

func handleGridCommands(parts []string) bool {
	if parts[0] != CmdShowGrid {
		return false
	}

	hasWrongArgCount := len(parts) < 2 || len(parts) > 3
	if hasWrongArgCount {
		fmt.Println(MsgInvalidShowGrid)
		return true
	}

	wIDStr := strings.TrimPrefix(parts[1], "W")
	wID, err := strconv.Atoi(wIDStr)
	isInvalidWarehouse := err != nil || !validWarehouseID(wID)
	if isInvalidWarehouse {
		fmt.Println(MsgInvalidWarehouseID)
		return true
	}

	isLiveMode := len(parts) == 3 && parts[2] == "live"
	if isLiveMode {
		fmt.Println("Live simulation mode. Press Ctrl+D to exit.")
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print(clearScreen)
			printGrid(wID)
			time.Sleep(1 * time.Second)

			if isEOF(reader) {
				break
			}
		}
		return true
	}

	printGrid(wID)
	return true
}

func printGrid(wID int) {
	warehouse := warehouses[wID-1]

	// Robot Grid
	robotGrid := [10][10]string{}
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			robotGrid[y][x] = MsgGridEmptyCell
		}
	}
	for i, r := range warehouse.Robots() {
		state := r.CurrentState()
		label := fmt.Sprintf("%s%d", MsgGridRobotPrefix, i+1)
		if strings.HasPrefix(fmt.Sprintf("%T", r), "*librobot.diagonalRobot") {
			label = fmt.Sprintf("%s%d", MsgGridDiagonalPrefix, i+1)
		}
		if state.X < 10 && state.Y < 10 {
			robotGrid[state.Y][state.X] = label
		}
	}
	robotGrid[0][0] = MsgGridOriginCell

	fmt.Println(MsgRobotGridHeader)
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			fmt.Printf("%-3s", robotGrid[y][x])
		}
		fmt.Println()
	}

	// Crate Grid
	if cw, ok := warehouse.(librobot.CrateWarehouse); ok {
		fmt.Printf("\n%s\n", MsgCrateGridHeader)
		crateGrid := [10][10]string{}
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				if cw.HasCrate(uint(x), uint(y)) {
					crateGrid[y][x] = MsgGridCrateCell
				} else {
					crateGrid[y][x] = MsgGridEmptyCell
				}
			}
		}
		crateGrid[0][0] = MsgGridOriginCell

		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				fmt.Printf("%-3s", crateGrid[y][x])
			}
			fmt.Println()
		}
	}
}

func isEOF(reader *bufio.Reader) bool {
	readerReady := false
	ch := make(chan struct{})
	go func() {
		_, err := reader.Peek(1)
		if err != nil {
			readerReady = true
		}
		close(ch)
	}()
	select {
	case <-ch:
		return readerReady
	case <-time.After(10 * time.Millisecond):
		return false
	}
}
