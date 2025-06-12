package librobot

import (
	"time"
)

const (
	WAREHOUSE_WIDTH  = 10
	WAREHOUSE_HEIGHT = 10
)

type MovementHandler func(r *robot, c string) (moved bool)

func handleNormal(r *robot, c string) bool {
	newX, newY := r.x, r.y
	switch c {
	case "N":
		if r.y > 0 {
			newY--
		} else {
			return false
		}
	case "S":
		if r.y < WAREHOUSE_HEIGHT-1 {
			newY++
		} else {
			return false
		}
	case "E":
		if r.x < WAREHOUSE_WIDTH-1 {
			newX++
		} else {
			return false
		}
	case "W":
		if r.x > 0 {
			newX--
		} else {
			return false
		}
	default:
		return false
	}

	if r.wh.IsOccupied(newX, newY) {
		return false
	}

	r.wh.UpdatePosition(r.x, r.y, newX, newY)
	r.x, r.y = newX, newY
	return true
}

func handleCrate(r *robot, c string) bool {
	switch c {
	case "G":
		if !r.hasCrate {
			if cw, ok := r.wh.(CrateWarehouse); ok && cw.HasCrate(r.x, r.y) {
				cw.DelCrate(r.x, r.y)
				r.hasCrate = true
				return true
			}
		}
	case "D":
		if r.hasCrate {
			if cw, ok := r.wh.(CrateWarehouse); ok && !cw.HasCrate(r.x, r.y) {
				cw.AddCrate(r.x, r.y)
				r.hasCrate = false
				return true
			}
		}
	}
	return false
}

func handleDiagonal(r *robot, c string) bool {
	newX, newY := r.x, r.y
	switch c {
	case "NE":
		newY--
		newX++
	case "NW":
		newY--
		newX--
	case "SE":
		newY++
		newX++
	case "SW":
		newY++
		newX--
	default:
		return false
	}

	if newX >= WAREHOUSE_WIDTH || newY >= WAREHOUSE_HEIGHT {
		return false
	}
	if r.wh.IsOccupied(newX, newY) {
		return false
	}

	r.wh.UpdatePosition(r.x, r.y, newX, newY)
	r.x, r.y = newX, newY
	return true
}

func runMovement(r *robot, commands []string, posCh chan RobotState, errCh chan error, handlers ...MovementHandler) {
	for _, c := range commands {
		time.Sleep(1 * time.Second)

		r.stepLock.Lock()
		handled := false
		for _, h := range handlers {
			if h(r, c) {
				handled = true
				break
			}
		}
		if handled {
			posCh <- RobotState{X: r.x, Y: r.y, HasCrate: r.hasCrate}
		}
		r.stepLock.Unlock()
	}
	errCh <- nil
	close(posCh)
}
