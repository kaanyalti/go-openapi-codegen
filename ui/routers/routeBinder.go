package routers

type RouteBinder interface {
	BindControllers(Router)
}
