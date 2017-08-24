package pool

import (
	"log"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/projects"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	resolvers["Project"] = projectsResolver
}

const (
	ProjectCreate   = Action("ProjectCreate")
	ProjectExists   = Action("ProjectExists")
	ProjectsAll     = Action("ProjectsAll")
	ProjectFindById = Action("ProjectFindById")
)

func projectsResolver(action Action) (service ServiceFunc) {
	switch action {
	case ProjectCreate:
		service = func(source db.DataSource, project interface{}) (interface{}, error) {
			return projects.CreateProject(source, project.(models.Project))
		}
		return

	case ProjectExists:
		service = func(source db.DataSource, project interface{}) (interface{}, error) {
			return projects.CheckProjectExists(source, project.(models.Project))
		}
		return

	case ProjectsAll:
		service = func(source db.DataSource, _ interface{}) (interface{}, error) {
			return projects.AllProjects(source)
		}
		return

	case ProjectFindById:
		service = func(source db.DataSource, id interface{}) (interface{}, error) {
			return projects.FindProjectById(source, id.(bson.ObjectId))
		}
		return
	default:
		log.Panicf("can not find resolver with action: %v, in projects resolvers", action)
		return
	}
}
