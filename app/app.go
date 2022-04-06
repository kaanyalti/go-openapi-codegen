package app

import (
	"net/http"
	"portfolioManagement/infrastructure"
	"portfolioManagement/router"
	utils "portfolioManagement/utils/logger"
)

type App struct {
	router router.Router
	logger utils.Logger
}

func NewApp() *App {
	return &App{
		router: router.NewRouter(),
		logger: ProvideLogger(),
	}
}

func (a *App) Start() {
	dbConn, err := infrastructure.SetupDbConnection(a.logger)
	if err != nil {
		panic(err)
	}

	db, err := dbConn.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := a.router.Start(":8080"); err != http.ErrServerClosed {
		a.logger.Error(err)
	}
}
