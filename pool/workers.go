package pool

import (
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/services"
	"io"
	"log"
	"runtime"
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
		result, err := services.GetUserByEmailAndPassword(mongo, job.ModelType)

		if err == io.EOF || err == io.ErrUnexpectedEOF {
			go func(job *Job, queue chan *Job) {
				queue <- job
			}(job, queue)

			mongo = connect()

			continue
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
