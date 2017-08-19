package pool

type WorkerServiceFunc func(model interface{}) interface{}
type workerServices map[string]WorkerServiceFunc

var workerServicesMap = make(map[string]WorkerServiceFunc)

func InitWorkersHandler() {
	for _, value := range typesList {
		workerServicesMap[value] = findServiceByAction(value)
	}
}

func findServiceByAction(actionType string) WorkerServiceFunc {
	//switch actionType {
	//	case
	//}
	return func(model interface{}) interface{} {
		return nil
	}
}
