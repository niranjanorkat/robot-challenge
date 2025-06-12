package repl

func validWarehouseID(wID int) bool {
	return wID > 0 && wID <= len(warehouses)
}
