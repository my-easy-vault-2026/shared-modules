package common

type WorkerPool int // 工作池
const (
	WORKER_POOL_DEFAULT WorkerPool = 1
)

func (wp WorkerPool) String() string {
	switch wp {
	case WORKER_POOL_DEFAULT:
		return "default"
	default:
		return ""
	}
}
