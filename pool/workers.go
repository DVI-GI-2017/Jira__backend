package pool

import (
	"fmt"
	"github.com/DVI-GI-2017/Jira__backend/auth"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"log"
	"runtime"
)

type Job struct {
	JobId int
	User  *auth.Credentials
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
		err := users.Insert(job.User)

		fmt.Println(err)

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

		user := new(auth.Credentials)

		err = users.Find(bson.M{"email": job.User.Email}).One(&user)

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

// test: for i in {1..15}; do echo '{"email": "test", "password": "password"}' | curl -d @- http://localhost:3000/api/v1/test; done
