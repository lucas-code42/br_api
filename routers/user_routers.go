package routers

import (
	"br_api/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartRouters() {
	r := mux.NewRouter()
	r.HandleFunc("/", api.GetAllAvailable)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)

	log.Fatal(http.ListenAndServe(":5000", r))

}
