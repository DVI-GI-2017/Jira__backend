package pool

import (
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/services"
	"gopkg.in/mgo.v2"
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
		conn := new(db.MongoConnection)
		conn.OriginalSession = mongo
		result, err := services.GetUserByEmailAndPassword(conn, job.ModelType)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			go func(job *Job, queue chan *Job) {
				queue <- job
			}(job, queue)
			mongo = connect(id)
			continue
		}

		results <- &JobResult{
			WorkerId:   id,
			Error:      err,
			ResultType: result,
		}
	}
}

func connect(workerId int) *mgo.Session {

	for {
		// Open a DB connection
		s, err := mgo.Dial("mongodb://localhost:27017/worker-test")
		if err != nil {
			log.Printf("Worker: Unable to connect to database (%s)", err)
			continue
		}

		// Connect to the DB collection
		return s

	}

}
