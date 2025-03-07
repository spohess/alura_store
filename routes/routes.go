package routes

import (
	"net/http"

	"github.com/spohess/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
