package repl

import (
	"fmt"
	"strings"
	"github.com/niranjanorkat/robot-challenge/librobot"
)

var warehouses []librobot.Warehouse

const (
	MsgWarehouseAdded   = "Warehouse added successfully."
	MsgNoWarehouses     = "No warehouses available."
	MsgCurrentWarehouses = "Current Warehouses:"
	MsgInvalidWarehouseID = "Invalid warehouse ID."
)

func handleWarehouseCommands(parts []string) bool {
	switch parts[0] {

	case CmdAddWarehouse:
		warehouseType := "n" // default to normal
		if len(parts) >= 2 {
			warehouseType = strings.ToLower(parts[1])
		}

		switch warehouseType {
		case "n", "":
			warehouses = append(warehouses, librobot.NewWarehouse())
			fmt.Println("Normal warehouse added.")
		case "c":
			warehouses = append(warehouses, librobot.NewCrateWarehouse())
			fmt.Println("Crate-enabled warehouse added.")
		default:
			fmt.Println("Unknown warehouse type. Use 'n' for normal or 'c' for crate-enabled.")
		}

	case CmdShowWarehouse:
		if len(warehouses) == 0 {
			fmt.Println(MsgNoWarehouses)
			return false
		}
		fmt.Println(MsgCurrentWarehouses)
		for i, wh := range warehouses {
			label := "Normal"
			if _, ok := wh.(librobot.CrateWarehouse); ok {
				label = "Crate-enabled"
			}
			fmt.Printf("- W%d (%s)\n", i+1, label)
		}

	default:
		return false
	}
	return true
}
