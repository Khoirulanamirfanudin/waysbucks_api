package routes

import (
	"be_waysbucks/handlers"
	"be_waysbucks/pkg/middleware"
	"be_waysbucks/pkg/mysql"
	"be_waysbucks/repositories"

	"github.com/gorilla/mux"
)

func ToppingRoutes(r *mux.Router) {
	toppingRepository := repositories.RepositoryTopping(mysql.DB)
	h := handlers.HandlerTopping(toppingRepository)

	r.HandleFunc("/toppings", middleware.Auth(h.FindToppings)).Methods("GET")
	r.HandleFunc("/topping/{id}", h.GetTopping).Methods("GET")
	// Create "/topping" route using middleware Auth, middleware UploadFile, handler Createtopping, and method POST
	r.HandleFunc("/topping", middleware.Auth(middleware.UploadFile(h.CreateTopping))).Methods("POST")
	r.HandleFunc("/topping/{id}", h.DeleteTopping).Methods("DELETE")
	r.HandleFunc("/topping/{id}", middleware.UploadFile(h.UpdateTopping)).Methods("PATCH")
}
