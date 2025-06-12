package librobot

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
			id:        id,
			x:         0,
			y:         0,
			wh:        wh,
			taskQueue: make(chan task, 100),
		},
	}
	r.taskProcessor()
	return r
}

// EnqueueTask processes commands with support for diagonal, normal, and crate movements.
func (r *diagonalRobot) EnqueueTask(commands string) (string, chan RobotState, chan error) {
	tokens := tokenizeCommands(commands)

	taskID := randomTaskID()
	posCh := make(chan RobotState)
	errCh := make(chan error, 1)

	r.taskQueue <- task{
		id:       taskID,
		commands: tokens,
		posCh:    posCh,
		errCh:    errCh,
	}

	return taskID, posCh, errCh
}

func (r *diagonalRobot) taskProcessor() {
	for t := range r.taskQueue {
		runMovement(r.robot, t.commands, t.posCh, t.errCh, handleNormal, handleCrate, handleDiagonal)
	}
}

func (r *diagonalRobot) CancelTask(taskID string) error {
	return r.robot.CancelTask(taskID)
}

func (r *diagonalRobot) CurrentState() RobotState {
	return r.robot.CurrentState()
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
