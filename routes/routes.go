package routes

func InitRouter(r *router) {
	loginRoutes(r)
	usersRoutes(r)
	testRoutes(r)
}
