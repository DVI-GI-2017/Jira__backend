package pool

type Job struct {
	JobId int
}

type JobResult struct {
	JobId      int
	WorkerId   int
	Error      error
	ResultType interface{}
}
