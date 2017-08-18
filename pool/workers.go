package pool

import "runtime"

type Job struct {
	JobId int
}

type JobResult struct {
	JobId      int
	WorkerId   int
	Error      error
	ResultType interface{}
}

func InitWorkers()  {
	for id := 0; id < runtime.NumCPU(); id++ {
		go worker(id, queue, results)
	}
}

func worker(id int, queue chan *Job, results chan<- *JobResult) {

}

