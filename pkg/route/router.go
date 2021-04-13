package route

import (
	"net/http"
	"github.com/gorilla/mux"
)

// func Initialize() {
//     Router = mux.NewRouter()
//     routes.RegisterWebRoutes(Router)
// }

// Name2URL 通过路由名称来获取 URL
func Name2URL(routeName string,pairs ...string) string  {
	var route *mux.Router
	url,err := route.Get(routeName).URL(pairs...)
	if err != nil {
		// checkError(err)
		return ""
	}
	return url.String()
}

// GetRouteVariable 获取 URI 路由参数
func GetRouterVariable(parameterName string,r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}