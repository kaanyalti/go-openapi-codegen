package routers

import "openapiCodegen/ui/controllers"

type PhotoRoutes struct {
	controllers controllers.PhotoControllers
}

func NewPhotoRoutes(controllers controllers.PhotoControllers) RouteBinder {
	return &PhotoRoutes{
		controllers: controllers,
	}
}

func (p *PhotoRoutes) BindControllers(router Router) {
	router.POST("/photos", p.controllers.CreatePhoto)
}
