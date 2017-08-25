package pool

import (
	"log"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/projects"
)

func init() {
	resolvers["Project"] = projectsResolver
}

const (
	ProjectCreate     = Action("ProjectCreate")
	ProjectExists     = Action("ProjectExists")
	ProjectsAll       = Action("ProjectsAll")
	ProjectFindById   = Action("ProjectFindById")
	ProjectAllUsers   = Action("ProjectAllUsers")
	ProjectAllTasks   = Action("ProjectAllTasks")
	ProjectAddUser    = Action("ProjectAddUser")
	ProjectDeleteUser = Action("ProjectDeleteUser")
	ProjectUserExists = Action("ProjectUserExists")
)

func projectsResolver(action Action) ServiceFunc {
	switch action {
	case ProjectCreate:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if project, ok := data.(models.Project); ok {
				return projects.CreateProject(source, project)
			}
			return models.Project{}, castFailsMsg(data, models.Project{})
		}

	case ProjectExists:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if project, ok := data.(models.Project); ok {
				return projects.CheckProjectExists(source, project)
			}
			return false, castFailsMsg(data, models.Project{})
		}

	case ProjectsAll:
		return func(source db.DataSource, _ interface{}) (interface{}, error) {
			return projects.AllProjects(source)
		}

	case ProjectFindById:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if id, ok := data.(models.RequiredId); ok {
				return projects.FindProjectById(source, id)
			}
			return models.Project{}, castFailsMsg(data, models.RequiredId{})
		}

	case ProjectAllUsers:
		return func(source db.DataSource, data interface{}) (result interface{}, err error) {
			if id, ok := data.(models.RequiredId); ok {
				return projects.AllUsersInProject(source, id)
			}
			return models.UsersList{}, castFailsMsg(data, models.RequiredId{})
		}

	case ProjectAllTasks:
		return func(source db.DataSource, data interface{}) (result interface{}, err error) {
			if id, ok := data.(models.RequiredId); ok {
				return projects.AllTasksInProject(source, id)
			}
			return models.TasksList{}, castFailsMsg(data, models.RequiredId{})
		}

	case ProjectAddUser:
		return func(source db.DataSource, data interface{}) (result interface{}, err error) {
			if ids, ok := data.(models.ProjectUser); ok {
				return projects.AddUserToProject(source, ids.ProjectId, ids.UserId)
			}
			return models.UsersList{}, castFailsMsg(data, models.ProjectUser{})
		}

	case ProjectDeleteUser:
		return func(source db.DataSource, data interface{}) (result interface{}, err error) {
			if ids, ok := data.(models.ProjectUser); ok {
				return projects.DeleteUserFromProject(source, ids.ProjectId, ids.UserId)
			}
			return models.UsersList{}, castFailsMsg(data, models.ProjectUser{})
		}

	case ProjectUserExists:
		return func(source db.DataSource, data interface{}) (result interface{}, err error) {
			if ids, ok := data.(models.ProjectUser); ok {
				return projects.CheckUserInProject(source, ids.UserId, ids.ProjectId)
			}
			return false, castFailsMsg(data, models.ProjectUser{})
		}

	default:
		log.Panicf("can not find resolver with action: %v, in projects resolvers", action)
		return nil
	}
}
