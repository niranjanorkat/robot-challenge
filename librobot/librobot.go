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
