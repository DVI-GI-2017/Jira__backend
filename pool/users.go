package pool

import (
	"log"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/users"
)

func init() {
	resolvers["User"] = usersResolver
}

const (
	UserCreate      = Action("UserCreate")
	UserExists      = Action("UserExists")
	UserAuthorize   = Action("UserAuthorize")
	UserFindById    = Action("UserFindById")
	UserFindByEmail = Action("UserFindByEmail")
	UsersAll        = Action("UsersAll")
	UserAllProjects = Action("UserAllProjects")
)

func usersResolver(action Action) ServiceFunc {
	switch action {

	case UserCreate:
		return func(source db.DataSource, data interface{}) (result interface{}, err error) {
			if user, ok := data.(models.User); ok {
				return users.CreateUser(source, user)
			}
			return models.User{}, castFailsMsg(data, models.User{})
		}

	case UserExists:
		return func(source db.DataSource, data interface{}) (result interface{}, err error) {
			if user, ok := data.(models.User); ok {
				return users.CheckUserExists(source, user)
			}
			return false, castFailsMsg(data, models.User{})
		}

	case UserAuthorize:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if user, ok := data.(models.User); ok {
				return users.CheckUserCredentials(source, user)
			}
			return false, castFailsMsg(data, models.User{})
		}

	case UsersAll:
		return func(source db.DataSource, _ interface{}) (result interface{}, err error) {
			return users.AllUsers(source)
		}

	case UserFindById:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if id, ok := data.(models.RequiredId); ok {
				return users.FindUserById(source, id)
			}
			return models.User{}, castFailsMsg(data, models.RequiredId{})
		}

	case UserFindByEmail:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if email, ok := data.(models.Email); ok {
				return users.FindUserByEmail(source, email)
			}
			return models.User{}, castFailsMsg(data, models.Email(""))
		}

	case UserAllProjects:
		return func(source db.DataSource, data interface{}) (result interface{}, err error) {
			if id, ok := data.(models.RequiredId); ok {
				return users.AllUserProjects(source, id)
			}
			return models.ProjectsList{}, castFailsMsg(data, models.RequiredId{})
		}

	default:
		log.Panicf("can not find resolver with action: %v, in users resolvers", action)
		return nil
	}
}
