package pool

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/users"
	"github.com/DVI-GI-2017/Jira__backend/utils"
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

func usersResolver(action Action) (ServiceFunc, error) {
	switch action {

	case UserCreate:
		return func(source db.DataSource, data interface{}) (result interface{}, err error) {
			if user, ok := data.(models.User); ok {
				return users.CreateUser(source, user)
			}
			return models.User{}, utils.ErrCastFails(data, models.User{})
		}, nil

	case UserExists:
		return func(source db.DataSource, data interface{}) (result interface{}, err error) {
			if user, ok := data.(models.User); ok {
				return users.CheckUserExists(source, user)
			}
			return false, utils.ErrCastFails(data, models.User{})
		}, nil

	case UserAuthorize:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if user, ok := data.(models.User); ok {
				return users.CheckUserCredentials(source, user)
			}
			return false, utils.ErrCastFails(data, models.User{})
		}, nil

	case UsersAll:
		return func(source db.DataSource, _ interface{}) (result interface{}, err error) {
			return users.AllUsers(source)
		}, nil

	case UserFindById:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if id, ok := data.(models.RequiredId); ok {
				return users.FindUserById(source, id)
			}
			return models.User{}, utils.ErrCastFails(data, models.RequiredId{})
		}, nil

	case UserFindByEmail:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if email, ok := data.(models.Email); ok {
				return users.FindUserByEmail(source, email)
			}
			return models.User{}, utils.ErrCastFails(data, models.Email(""))
		}, nil

	case UserAllProjects:
		return func(source db.DataSource, data interface{}) (result interface{}, err error) {
			if id, ok := data.(models.RequiredId); ok {
				return users.AllUserProjects(source, id)
			}
			return models.ProjectsList{}, utils.ErrCastFails(data, models.RequiredId{})
		}, nil

	default:
		return nil, fmt.Errorf("can not find resolver with action: %v, in users resolvers", action)

	}
}
