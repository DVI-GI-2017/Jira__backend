package pool

type Action string

const (
	//Users actions
	CreateUser           = Action("CreateUser")
	CheckUserExists      = Action("CheckUserExists")
	CheckUserCredentials = Action("CheckUserCredentials")
	FindUserById         = Action("FindUserById")
	AllUsers             = Action("AllUsers")

	// Projects actions
	CreateProject      = Action("CreateProject")
	CheckProjectExists = Action("CheckProjectExists")
	AllProjects        = Action("AllProjects")
	FindProjectById    = Action("FindProjectById")

	// Tasks actions
	CreateTask      = Action("CreateTask")
	CheckTaskExists = Action("CheckTaskExists")
	AllTasks        = Action("AllTasks")
	FindTaskById    = Action("FindTaskById")

	// Labels actions
	AddLabelToTask       = Action("AddLabelToTask")
	AllLabelsOnTask      = Action("AllLabelsOnTask")
	CheckLabelAlreadySet = Action("CheckLabelAlreadySet")
	DeleteLabelFromTask  = Action("DeleteLabelFromTask")
)
