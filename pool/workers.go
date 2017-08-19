package pool

import (
	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
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
	mongo := connect(id)

	for job := range queue {
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			go func(job *Job, queue chan *Job) {
				queue <- job
			}(job, queue)
			mongo = connect(id)
			continue
		}

		user := new(models.User)

		err = users.Find(nil).One(&user)

		results <- &JobResult{
			WorkerId:   id,
			Error:      err,
			ResultType: user,
		}
	}
}

func connect(workerId int) *db.MongoConnection {
	for {
		log.Printf("Worker %d: Connecting to db", workerId)

		newConnection, err := db.NewDBConnection(configs.ConfigInfo.Mongo)
		if err != nil || newConnection == nil {
			log.Panicf("can not start db: %s", err)

			continue
		}

		return newConnection
	}
}
