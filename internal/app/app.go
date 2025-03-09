package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tanvir-rifat007/workoutApi/internal/api"
)


type Application struct{
	Logger *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error){


	return &Application{
		Logger: log.New(os.Stdout,"",log.Ldate | log.Ltime),
		WorkoutHandler: api.CreateNewWorkoutHander(),


	}, nil
}


// controller:
func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w,"Health Check\n")
}

