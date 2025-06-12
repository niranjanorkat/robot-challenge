package repl

import (
	"fmt"
	"strings"
	"github.com/niranjanorkat/robot-challenge/librobot"
)

var warehouses []librobot.Warehouse

const (
	MsgUsageAddWarehouse = "add_warehouse [type] - Adds a warehouse. Type: 'n' (normal) or 'c' (crate-enabled). Defaults to 'n'."
	MsgUsageShowWarehouse = "show_warehouse - Lists all warehouses with their type."
	MsgWarehouseAdded   = "Warehouse added successfully."
	MsgNoWarehouses     = "No warehouses available."
	MsgCurrentWarehouses = "Current Warehouses:"
	MsgInvalidWarehouseID = "Invalid warehouse ID."
	MsgNormalWarehouseAdded = "Normal warehouse added."
	MsgCrateWarehouseAdded = "Crate-enabled warehouse added."
	MsgUnknownWarehouseType = "Unknown warehouse type. Use 'n' for normal or 'c' for crate-enabled."
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
			fmt.Println(MsgNormalWarehouseAdded)
		case "c":
			warehouses = append(warehouses, librobot.NewCrateWarehouse())
			fmt.Println(MsgCrateWarehouseAdded)
		default:
			fmt.Println(MsgUnknownWarehouseType)
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

func validWarehouseID(wID int) bool {
	return wID > 0 && wID <= len(warehouses)
}
