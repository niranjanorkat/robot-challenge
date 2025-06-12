package librobot
import (
	"sync"
)
type DiagonalRobot interface {
	Robot
}

type diagonalRobot struct {
	*robot
}

// NewDiagonalRobot constructs a diagonal-capable robot at position (0,0)
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

// EnqueueTask processes commands with support for diagonal, normal, and crate movements.
func (r *diagonalRobot) EnqueueTask(commands string) string {
	tokens := tokenizeCommands(commands)
	taskID := randomTaskID()

	t := &task{
		id:         taskID,
		commands:   tokens,
		rawCommand: commands,
		status:     TaskStatusOngoing,
		stop:       make(chan struct{}, 1),
	}

	r.activeTasks.Store(taskID, t)
	r.taskQueue <- *t
	return taskID
}

func (r *diagonalRobot) taskProcessor() {
	for t := range r.taskQueue {
		status := runMovement(r.robot, t.commands, t.stop, handleNormal, handleCrate, handleDiagonal)
		t.status = status
		r.activeTasks.Delete(t.id)
	}
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

func tokenizeCommands(input string) []string {
	result := []string{}
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
