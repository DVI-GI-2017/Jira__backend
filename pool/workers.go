package pool

import (
	"runtime"

	"github.com/DVI-GI-2017/Jira__backend/db"
)

// Type and channel for jobs
type job struct {
	service ServiceFunc
	input   interface{}
}

// Process job with self-contained input and given data source
func (j job) process() {
	source := db.Copy()
	defer source.Close()

	result, err := j.service(source, j.input)

	results <- &jobResult{
		err:    err,
		result: result,
	}
}

var jobs = make(chan *job, 512)

// Type and channel for results
type jobResult struct {
	err    error
	result interface{}
}

var results = make(chan *jobResult, 512)

func InitWorkers() {
	for id := 0; id < runtime.NumCPU(); id++ {
		go worker()
	}
}

func worker() {
	for job := range jobs {
		job.process()
	}
}
