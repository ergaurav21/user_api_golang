package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"user_api/controller"
	"user_api/cors"
	"user_api/database"
)

func  indexHandler(w http.ResponseWriter, r *http.Request){
	log.Println(r.Method)
	w.Write([]byte("<h1>Welcome to my web service</h1>"))
}

func main() {
	err := database.SetupDatabase()
	if err != nil {
		panic(err)
	}
    u :=   controller.NewUserHandler("hello")

    http.Handle("/users",cors.Middleware(u))
	http.Handle("/users/",cors.Middleware(u))

    http.HandleFunc("/",indexHandler)

    log.Println("server started...")
	log.Fatal(http.ListenAndServe(":5000",nil))

}

