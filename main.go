package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/tanvir-rifat007/workoutApi/internal/app"
	"github.com/tanvir-rifat007/workoutApi/internal/routes"
)

func main(){

	var port int;
	flag.IntVar(&port,"port",8080,"Port to run the server on")
	flag.Parse()

	app,err:= app.NewApplication()


	if err!=nil{
		panic(err)

	}



	// Create a new server
	server:= &http.Server{
		Addr: fmt.Sprintf(":%d",port),
		Handler: routes.SetupRoutes(app),
		IdleTimeout: time.Minute,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 30*time.Second,
	}

	fmt.Printf("Server is running on port %d\n",port)

	err = server.ListenAndServe()



	if err!=nil{
		app.Logger.Fatal(err)
	}
}


func healthCheck(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Health Check\n")
}