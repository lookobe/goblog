package bootstrap

import (
	"goblog/routes"
	"github.com/gorilla/mux"
	"goblog/pkg/route"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)
	route.SetRoute(router)
	return router
}