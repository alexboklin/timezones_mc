package main

import (
	"github.com/bapjiws/local_times_dashboard_backend/app/handlers"
	"github.com/bapjiws/local_times_dashboard_backend/app/middleware"
	"github.com/bapjiws/local_times_dashboard_backend/datastore/elasticsearch"
	"github.com/bapjiws/local_times_dashboard_backend/datastore/elasticsearch/configs"
	"github.com/gin-gonic/gin"
)

const (
	API_BASE = "api"
	PORT     = ":8888"
)

var context middleware.Context

func init() {
	context = middleware.Context{
		DS: elasticsearch.NewElasticStore(configs.CityStoreConfig),
	}
}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.Use(
		middleware.SetContext(context),
		middleware.AllowCors(),
	)

	cityRouter := router.Group(API_BASE)
	cityRouter.GET("/city", handlers.SuggestCities)
	cityRouter.GET("/city/:id", handlers.FindCityById)

	router.Run(PORT)
}
