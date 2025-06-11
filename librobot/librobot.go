package librobot

type command byte

const (
	North command = 'N'
	South command = 'S'
	East  command = 'E'
	West  command = 'W'
	Grab  command = 'G'
	Drop  command = 'D'
)

// CrateWarehouse provides an abstraction of a simulated warehouse containing both robots and crates.
type CrateWarehouse interface {
	Warehouse

	AddCrate(x uint, y uint) error
	DelCrate(x uint, y uint) error
}

