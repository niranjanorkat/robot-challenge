package librobot

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
)

type Robot interface {
	EnqueueTask(commands string) (taskID string)
	CancelTask(taskID string) error
	CurrentState() RobotState
}

type RobotState struct {
	X        uint
	Y        uint
	HasCrate bool
}

type robot struct {
	id        string
	x, y      uint
	hasCrate  bool
	wh        Warehouse
	taskQueue chan task
	stepLock  sync.Mutex
}

type task struct {
	id       string
	commands []string
	posCh    chan RobotState
	errCh    chan error
}

func NewRobot(id string, wh Warehouse) Robot {
	r := &robot{
		id:        id,
		x:         0,
		y:         0,
		wh:        wh,
		taskQueue: make(chan task, 100),
	}
	go r.taskProcessor()
	return r
}

func (r *robot) EnqueueTask(commands string) (string) {
	tokens := make([]string, len(commands))
	for i, c := range commands {
		tokens[i] = string(c)
	}

	taskID := randomTaskID()

	r.taskQueue <- task{
		id:       taskID,
		commands: tokens,
	}

	return taskID
}

func (r *robot) taskProcessor() {
	for t := range r.taskQueue {
		runMovement(r, t.commands, handleNormal, handleCrate)
	}
}

func (r *robot) CancelTask(taskID string) error {
	panic("CancelTask not implemented")
}

func (r *robot) CurrentState() RobotState {
	return RobotState{X: r.x, Y: r.y, HasCrate: r.hasCrate}
}

func randomTaskID() string {
	b := make([]byte, 4)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
