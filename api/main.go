package main

import (
	"log"
	"net/http"
	"os"

	"github.com/oluwafenyi/go-api-project-template/db"
	"github.com/oluwafenyi/go-api-project-template/routes"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	routes.Router.Mount("/", routes.GeneralRoutes())
	routes.Router.Mount("/user", routes.UserRoutes())

	log.Println("server listening at port 8000")
	err := http.ListenAndServe(":8000", routes.Router)
	if err != nil {
		db.DB.Close()
		log.Fatalln(err)
	}
}
