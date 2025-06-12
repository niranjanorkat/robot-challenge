package repl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/niranjanorkat/robot-challenge/librobot"
)

const (
	MsgUsageAddCrate = "add_crate W<id> x y - Adds a crate to (x, y) in a crate-enabled warehouse."
	MsgUsageShowCrates = "show_crates W<id> - Lists all crate locations in the warehouse."
	MsgCrateAdded       = "Crate added successfully."
	MsgCrateError       = "Error adding crate:"
	MsgCrateListHeader  = "Crates in warehouse:"
	MsgNoCrates         = "No crates in this warehouse."
	MsgNotCrateWarehouse = "This warehouse does not support crates."
	MsgCrateLocation    = "- Crate at (x=%d, y=%d)\n"
)

var (
	MsgInvalidAddCrate  = "Invalid add_crate command. Usage: " + MsgUsageAddCrate
	MsgInvalidShowCrate = "Invalid show_crates command. Usage: " + MsgUsageShowCrates
)

func handleCrateCommands(parts []string) bool {
	switch parts[0] {
	case CmdAddCrate:
		if len(parts) != 4 {
			fmt.Println(MsgInvalidAddCrate)
			return true
		}
		wID, err1 := strconv.Atoi(strings.TrimPrefix(parts[1], "W"))
		x, err2 := strconv.Atoi(parts[2])
		y, err3 := strconv.Atoi(parts[3])
		if err1 != nil || err2 != nil || err3 != nil || !validWarehouseID(wID) {
			fmt.Println(MsgInvalidAddCrate)
			return true
		}
		if cw, ok := warehouses[wID-1].(librobot.CrateWarehouse); ok {
			err := cw.AddCrate(uint(x), uint(y))
			if err != nil {
				fmt.Println(MsgCrateError, err)
			} else {
				fmt.Println(MsgCrateAdded)
			}
		} else {
			fmt.Println(MsgNotCrateWarehouse)
		}
		return true

	case CmdShowCrates:
		if len(parts) != 2 {
			fmt.Println(MsgInvalidShowCrate)
			return true
		}
		wID, err := strconv.Atoi(strings.TrimPrefix(parts[1], "W"))
		if err != nil || !validWarehouseID(wID) {
			fmt.Println(MsgInvalidWarehouseID)
			return true
		}
		cw, ok := warehouses[wID-1].(librobot.CrateWarehouse)
		if !ok {
			fmt.Println(MsgNotCrateWarehouse)
			return true
		}
		fmt.Println(MsgCrateListHeader)
		found := false
		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				if cw.HasCrate(uint(x), uint(y)) {
					fmt.Printf(MsgCrateLocation, x, y)
					found = true
				}
			}
		}
		if !found {
			fmt.Println(MsgNoCrates)
		}
		return true

	default:
		return false
	}
}
