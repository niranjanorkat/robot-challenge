package librobot

import (
	"sync"
)

// DiagonalRobot extends the Robot interface to support diagonal movement.
// It behaves like a regular Robot but can also interpret and execute diagonal commands like "NE", "NW", "SE", and "SW".
type DiagonalRobot interface {
	Robot
	IsDiagonal() bool
}

type diagonalRobot struct {
	*robot
}

// NewDiagonalRobot creates a new robot capable of diagonal movement.
func NewDiagonalRobot(id string, wh Warehouse) Robot {
	r := &diagonalRobot{
		robot: &robot{
			id:          id,
			x:           0,
			y:           0,
			wh:          wh,
			taskQueue:   make(chan task, 100),
			activeTasks: sync.Map{},
		},
	}
	go r.taskProcessor()
	return r
}

func (r *diagonalRobot) EnqueueTask(commands string) string {
	tokens := tokenizeDiagonalMoves(commands)
	taskID := generateTaskID()

	t := &task{
		id:         taskID,
		commands:   tokens,
		rawCommand: commands,
		status:     TaskStatusOngoing,
		stop:       make(chan struct{}),
	}

	r.activeTasks.Store(taskID, t)
	r.taskQueue <- *t
	return taskID
}

func (r *diagonalRobot) CancelTask(taskID string) error {
	return r.robot.CancelTask(taskID)
}

func (r *diagonalRobot) CurrentState() RobotState {
	return r.robot.CurrentState()
}

func (r *diagonalRobot) GetActiveTasks() []TaskInfo {
	return r.robot.GetActiveTasks()
}

func (r *diagonalRobot) IsDiagonal() bool {
	return true
}

func (r *diagonalRobot) taskProcessor() {
	for t := range r.taskQueue {
		status := runMovement(r.robot, t.commands, t.stop, handleNormal, handleCrate, handleDiagonal)
		t.status = status
		r.activeTasks.Delete(t.id)
	}
}

func tokenizeDiagonalMoves(input string) []string {
	var result []string
	i := 0
	for i < len(input) {
		if i+1 < len(input) {
			pair := input[i : i+2]
			switch pair {
			case "NE", "NW", "SE", "SW":
				result = append(result, pair)
				i += 2
				continue
			}
		}
		result = append(result, string(input[i]))
		i++
	}
	return result
}
