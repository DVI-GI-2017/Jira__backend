package pool

import (
	"Jira__backend/models"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"log"
	"runtime"
)

type Job struct {
	JobId int
	Email string
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

		// Perform the database query
		err := users.Insert(models.User{
			FirstName: fmt.Sprintf("User %d", "Vasya"),
			Email:     fmt.Sprintf("user-%s@example.com", job.Email),
		})

		if err == io.EOF || err == io.ErrUnexpectedEOF {
			// Our job hasn't completed because the database is no longer connected
			// Put our job back onto the queue (in another go routine to avoid blocking if queue buffer is full)
			// Then reconnect the database and continue processing
			go func(job *Job, queue chan *Job) {
				queue <- job
			}(job, queue)
			users = connect(id, session)
			continue
		}

		user := new(models.User)
		userid := fmt.Sprintf("User %s", job.Email)

		err = users.Find(bson.M{
			"$and": []interface{}{
				bson.M{"email": userid},
			},
		}).One(&user)

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
