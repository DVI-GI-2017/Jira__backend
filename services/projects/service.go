package projects

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2/bson"
)

const cProjects = "projects"

// Check if project with title == project.Title exists
func CheckProjectExists(source db.DataSource, project models.Project) (bool, error) {
	empty, err := source.C(cProjects).Find(bson.M{"title": project.Title}).IsEmpty()
	if err != nil {
		return false, fmt.Errorf("can not check if project '%v' exists: %v", project, err)
	}
	return !empty, err
}

// Creates project and returns it.
func CreateProject(source db.DataSource, project models.Project) (models.Project, error) {
	project.Id = bson.NewObjectId()

	err := source.C(cProjects).Insert(project)
	if err != nil {
		return models.Project{}, fmt.Errorf("can not create project '%v': %v", project, err)
	}
	return project, nil
}

// Returns all projects.
func AllProjects(source db.DataSource) (models.ProjectsList, error) {
	result, err := source.C(cProjects).Find(bson.M{}).All()
	if err != nil {
		return models.ProjectsList{}, fmt.Errorf("can not retrieve all projects: %v", err)
	}
	return result.(models.ProjectsList), nil
}

// Returns task with given id.
func FindProjectById(mongo db.DataSource, id bson.ObjectId) (models.Project, error) {
	result, err := mongo.C(cProjects).FindId(id).One()
	if err != nil {
		return models.Project{}, fmt.Errorf("can not find project with id '%s': %v", id, err)
	}
	return result.(models.Project), nil
}
