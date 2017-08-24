package pool

import "log"

func getServiceByAction(action Action) (service ServiceFunc, err error) {
	for prefix, resolver := range resolvers {
		if action.HasPrefix(prefix) {
			return resolver(action)
		}
	}
	log.Panicf("can not resolve service by action: %v", action)
	return
}

// Creates job with given action and input and returns result.
func Dispatch(action Action, input interface{}) (result interface{}, err error) {
	Queue <- &Job{
		Input:  input,
		Action: action,
	}

	jobResult := <-Results

	result = jobResult.Result
	err = jobResult.Error

	return
}
