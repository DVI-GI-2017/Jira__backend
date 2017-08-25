package pool

import "github.com/DVI-GI-2017/Jira__backend/db"

type Job struct {
	service ServiceFunc
	input   interface{}
}

// Process job with self-contained input and given data source
func (j Job) process(source db.DataSource) (result interface{}, err error) {
	return j.service(source, j.input)
}

type JobResult struct {
	workerId int
	err      error
	result   interface{}
}
