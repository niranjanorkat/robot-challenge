package librobot
import "github.com/google/uuid"
const (
	TaskStatusOngoing   = "ongoing"
	TaskStatusSuccessful = "successful"
	TaskStatusAborted   = "aborted"
)

// TaskInfo represents metadata about a task, used for reporting active tasks.
type TaskInfo struct {
	ID          string
	Status      string
	RawCommand  string
}

type task struct {
	id         string
	commands   []string
	rawCommand string
	status     string
	stop       chan struct{}
}

func generateTaskID() string {
	return uuid.NewString()
}