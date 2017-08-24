package pool

type Action string

const (
	//Users actions
	UserCreate    = Action("UserCreate")
	UserExists    = Action("UserExists")
	UserAuthorize = Action("UserAuthorize")
	UserFindById  = Action("UserFindById")
	UsersAll      = Action("UsersAll")

	// Projects actions
	ProjectCreate   = Action("ProjectCreate")
	ProjectExists   = Action("ProjectExists")
	ProjectsAll     = Action("ProjectsAll")
	ProjectFindById = Action("ProjectFindById")

	// Tasks actions
	TaskCreate   = Action("TaskCreate")
	TaskExists   = Action("TaskExists")
	TasksAll     = Action("TasksAll")
	TaskFindById = Action("TaskFindById")

	// Labels actions
	LabelAddToTask      = Action("LabelAddToTask")
	LabelsAllOnTask     = Action("LabelsAllOnTask")
	LabelAlreadySet     = Action("LabelAlreadySet")
	LabelDeleteFromTask = Action("LabelDeleteFromTask")
)
