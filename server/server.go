package server

import (
	"net/http"
	"openapiCodegen/application"
	"openapiCodegen/infrastructure"
	"openapiCodegen/ui/controllers"
	"openapiCodegen/ui/routers"
	"openapiCodegen/utils"
)

type Server struct {
	config utils.Config
	router routers.Router
	logger utils.Logger
}

func NewServer(config utils.Config, router routers.Router, logger utils.Logger) Server {
	return Server{
		config: config,
		router: router,
		logger: logger,
	}
}

func (s Server) Serve() {
	dbConn, err := infrastructure.SetupDbConnection(s.config, s.logger)
	if err != nil {
		s.logger.Fatal(err)
	}

	db, err := dbConn.DB()
	if err != nil {
		s.logger.Fatal(err)
	}
	defer db.Close()

	photoStorage := infrastructure.NewGcsPhotoStorage()
	photoRepository := infrastructure.NewPostgresPhotoRepository(dbConn)
	photoService := application.NewPhotoService(photoRepository, photoStorage)
	photoControllers := controllers.NewPhotoControllers(photoService)
	photoRoutes := routers.NewPhotoRoutes(*photoControllers)
	photoRoutes.BindControllers(s.router)

	if err := s.router.Start(":8080"); err != http.ErrServerClosed {
		s.logger.Fatal(err)
	}
}
