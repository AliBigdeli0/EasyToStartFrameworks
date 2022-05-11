package worker

const (
	KILL_ALL = 1
)

type Job struct {
	Id      int
	Args    interface{}
	Returns interface{}
}
