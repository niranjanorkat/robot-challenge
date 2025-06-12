package librobot
import "github.com/google/uuid"
const (
	TaskStatusOngoing   = "ongoing"
	TaskStatusSuccessful = "successful"
	TaskStatusAborted   = "aborted"
)

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