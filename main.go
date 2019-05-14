package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arithmetics/got_boardgame/app"

	"github.com/arithmetics/got_boardgame/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/me", controllers.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET") //  user/2/contacts
	router.HandleFunc("/api/test", controllers.Test).Methods("GET")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)
	//cors optionsGoes Below
	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"*"})
	credsOk := handlers.AllowCredentials()

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk, credsOk)(router)))

}
