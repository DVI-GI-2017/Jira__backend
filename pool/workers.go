package pool

import (
	"io"
	"runtime"

	"github.com/DVI-GI-2017/Jira__backend/db"
)

type Job struct {
	Action *Action
	Input  interface{}
}

type JobResult struct {
	WorkerId   int
	Error      error
	Result interface{}
}

var Queue = make(chan *Job, 512)
var Results = make(chan *JobResult, 512)

func InitWorkers() {
	for id := 0; id < runtime.NumCPU(); id++ {
		go worker(id, Queue, Results)
	}
}

func worker(id int, queue chan *Job, results chan<- *JobResult) {
	for job := range queue {
		function, _ := GetServiceByAction(job.Action)

		result, err := function(db.GetDB(), job.Input)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			go func(job *Job, queue chan *Job) {
				queue <- job
			}(job, queue)

			continue
		}

		results <- &JobResult{
			WorkerId:   id,
			Error:      err,
			Result: result,
		}
	}
}
