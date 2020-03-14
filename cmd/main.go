package main

import (
	"fmt"
	"github.com/dnlgrwd87/blog-api/config"
	"github.com/dnlgrwd87/blog-api/users"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

func main() {
	db, err := config.InitializeDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	config.Migrate(db)
	config.AddContstraints(db)

	userService := users.UserService{DB: db}
	userHandler := users.UserHandler{Service: userService}

	r := chi.NewRouter()

	r.Use(config.GetCors().Handler)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.GetUsers)
		r.Get("/{id}", userHandler.GetUserById)
		r.Post("/", userHandler.CreateUser)
	})

	fmt.Println("Listenting on port 3000")

	http.ListenAndServe(":3000", r)
}
