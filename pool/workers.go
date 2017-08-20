package pool

import (
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/tools"
	"io"
	"log"
	"runtime"
	"strings"
)

type Job struct {
	Action    *Action
	ModelType interface{}
}

type JobResult struct {
	WorkerId   int
	Error      error
	ResultType interface{}
}

var Queue = make(chan *Job, 512)
var Results = make(chan *JobResult, 512)

func InitWorkers() {
	for id := 0; id < runtime.NumCPU(); id++ {
		go worker(id, Queue, Results)
	}
}

func worker(id int, queue chan *Job, results chan<- *JobResult) {
	mongo := connect()

	for job := range queue {
		function, _ := GetServiceByAction(job.Action)

		result, err := function(mongo, job.ModelType)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			go func(job *Job, queue chan *Job) {
				queue <- job
			}(job, queue)

			mongo = connect()

			continue
		}

		if err.Error() == "not found" &&
			strings.ToLower(tools.GetType(job.ModelType)) == "user" {
		}

		results <- &JobResult{
			WorkerId:   id,
			Error:      err,
			ResultType: result,
		}
	}
}

func connect() *db.MongoConnection {
	for {
		mongo, err := db.NewDBConnection()
		if err != nil {
			log.Printf("Worker: Unable to connect to database (%s)", err)
			continue
		}

		return mongo
	}
}
