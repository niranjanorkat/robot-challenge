package repl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/niranjanorkat/robot-challenge/librobot"
)

const (
	MsgInvalidShowGrid = "Usage: show_grid W<id>"
)

func handleGridCommands(parts []string) bool {
	if parts[0] != CmdShowGrid {
		return false
	}
	if len(parts) != 2 {
		fmt.Println(MsgInvalidShowGrid)
		return true
	}
	wID, err := strconv.Atoi(strings.TrimPrefix(parts[1], "W"))
	if err != nil || !validWarehouseID(wID) {
		fmt.Println(MsgInvalidWarehouseID)
		return true
	}
	warehouse := warehouses[wID-1]

	// Robot Grid
	robotGrid := [10][10]string{}
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			robotGrid[y][x] = "."
		}
	}
	for i, r := range warehouse.Robots() {
		state := r.CurrentState()
		label := fmt.Sprintf("R%d", i+1)
		if strings.HasPrefix(fmt.Sprintf("%T", r), "*librobot.diagonalRobot") {
			label = fmt.Sprintf("D%d", i+1)
		}
		if state.X < 10 && state.Y < 10 {
			robotGrid[state.Y][state.X] = label
		}
	}
	robotGrid[0][0] = "O"

	// Print robot grid first
	fmt.Println("Robot Grid")
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			fmt.Printf("%-3s", robotGrid[y][x])
		}
		fmt.Println()
	}

	// Only show crate grid if it's a crate warehouse
	if cw, ok := warehouse.(librobot.CrateWarehouse); ok {
		fmt.Println("\nCrate Grid")
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
		crateGrid[0][0] = "O"

		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				fmt.Printf("%-3s", crateGrid[y][x])
			}
			fmt.Println()
		}
	}

	return true
}
