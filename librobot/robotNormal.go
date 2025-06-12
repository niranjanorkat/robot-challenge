package librobot

import (
	"sync"
	"fmt"
)

type Robot interface {
	EnqueueTask(commands string) (taskID string)
	CancelTask(taskID string) error
	CurrentState() RobotState
	GetActiveTasks() []TaskInfo
}

type RobotState struct {
	X        uint
	Y        uint
}

type robot struct {
	id          string
	x, y        uint
	isCarryingCrate    bool
	wh          Warehouse
	taskQueue   chan task
	stepLock    sync.Mutex
	activeTasks sync.Map
}

// ─── Constructor ─────────────────────────────────────────────

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

// ─── Public Methods ──────────────────────────────────────────

func (r *robot) EnqueueTask(commands string) string {
	tokens := make([]string, len(commands))
	for i, c := range commands {
		tokens[i] = string(c)
	}

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

func (r *robot) CancelTask(taskID string) error {
	if val, ok := r.activeTasks.Load(taskID); ok {
		t := val.(*task)
		select {
		case t.stop <- struct{}{}:
			t.status = TaskStatusAborted
			r.activeTasks.Delete(taskID)
			return nil
		default:
			return fmt.Errorf("task %s already completed or cancelling", taskID)
		}
	}
	return fmt.Errorf("task %s not found", taskID)
}

func (r *robot) GetActiveTasks() []TaskInfo {
	var infos []TaskInfo
	r.activeTasks.Range(func(_, v any) bool {
		t := v.(*task)
		infos = append(infos, TaskInfo{
			ID:         t.id,
			Status:     t.status,
			RawCommand: t.rawCommand,
		})
		return true
	})
	return infos
}

func (r *robot) CurrentState() RobotState {
	return RobotState{X: r.x, Y: r.y}
}

// ─── Internal Processor ──────────────────────────────────────

func (r *robot) taskProcessor() {
	for t := range r.taskQueue {
		status := runMovement(r, t.commands, t.stop, handleNormal, handleCrate)
		t.status = status
		r.activeTasks.Delete(t.id)
	}
}
