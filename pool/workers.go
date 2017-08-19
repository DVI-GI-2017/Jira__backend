package pool

import (
	"fmt"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2"
	"io"
	"log"
	"runtime"
)

type Job struct {
	JobId     int
	Action    string
	ModelType interface{}
}

type JobResult struct {
	JobId      int
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
	var session *mgo.Session
	users := connect(id, session)

	for job := range queue {

		fmt.Println("Data:")
		fmt.Println(job.ModelType)
		err := users.Insert(job.ModelType)

		if err == io.EOF || err == io.ErrUnexpectedEOF {
			go func(job *Job, queue chan *Job) {
				queue <- job
			}(job, queue)
			users = connect(id, session)
			continue
		}

		user := new(models.User)

		err = users.Find(nil).One(&user)

		// Send our results back
		results <- &JobResult{
			JobId:      job.JobId,
			WorkerId:   id,
			Error:      err,
			ResultType: user,
		}
	}
}

func connect(workerId int, session *mgo.Session) *mgo.Collection {
	for {
		// Open a DB connection
		log.Printf("Worker %d: Connecting to", workerId)
		s, err := mgo.Dial("mongodb://localhost:27017/worker-test")
		if err != nil {
			log.Printf("Worker %d: Unable to connect to database (%s)", workerId, err)
			continue
		}

		// Connect to the DB collection
		return s.DB("worker-test").C("users")
	}
}
