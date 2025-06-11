package librobot

import (
	"time"
)
const (
	WAREHOUSE_WIDTH  = 10
	WAREHOUSE_HEIGHT = 10
)
// Robot provides an abstraction of a warehouse robot which accepts tasks in the form of strings of commands.
type Robot interface {
	EnqueueTask(commands string) (taskID string, position chan RobotState, err chan error)
	CancelTask(taskID string) error
	CurrentState() RobotState
}

// RobotState provides an abstraction of the state of a warehouse robot.
type RobotState struct {
	X uint
	Y uint
	HasCrate bool
}

type robot struct {
	id string
	x, y uint
	wh *warehouse
}

// NewRobot constructs a new robot at a given position.
func NewRobot(id string, wh *warehouse) Robot {
	return &robot{
		id: id,
		x:  0,
		y:  0,
		wh: wh,
	}
}


func (r *robot) EnqueueTask(commands string) (string, chan RobotState, chan error) {
	positionCh := make(chan RobotState)
	errCh := make(chan error, 1)
	hasCrate := false

	go func() {
		for _, c := range commands {
			time.Sleep(1 * time.Second)
			newX, newY := r.x, r.y

			switch c {
			case 'N':
				if r.y > 0 {
					newY--
				}
			case 'S':
				if r.y < WAREHOUSE_HEIGHT-1 {
					newY++
				}
			case 'E':
				if r.x < WAREHOUSE_WIDTH-1 {
					newX++
				}
			case 'W':
				if r.x > 0 {
					newX--
				}
			case 'G':
				if !hasCrate {
					if r.wh.HasCrate(r.x, r.y) {
						r.wh.DelCrate(r.x, r.y)
						hasCrate = true
					}
				}
			case 'D':
				if hasCrate {
					if !r.wh.HasCrate(r.x, r.y) {
						r.wh.AddCrate(r.x, r.y)
						hasCrate = false
					}
				}
			}

			if c == 'N' || c == 'S' || c == 'E' || c == 'W' {
				if r.wh.IsOccupied(newX, newY) {
					continue
				}
				r.wh.updatePosition(r.x, r.y, newX, newY)
				r.x, r.y = newX, newY
			}

			positionCh <- RobotState{X: r.x, Y: r.y, HasCrate: hasCrate}
		}
		close(positionCh)
		errCh <- nil
	}()

	return "task-id", positionCh, errCh
}

func (r *robot) CancelTask(taskID string) error {
	panic("CancelTask not implemented")
}

func (r *robot) CurrentState() RobotState {
	return RobotState{X: r.x, Y: r.y}
}
