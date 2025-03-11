package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tanvir-rifat007/workoutApi/internal/api"
	"github.com/tanvir-rifat007/workoutApi/internal/store"
	"github.com/tanvir-rifat007/workoutApi/migrations"
)


type Application struct{
	Logger *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB *sql.DB
}

func NewApplication() (*Application, error){
	db,err:= store.Open()

	if err!=nil{
		return nil,err
	}

	err = store.MigrateFS(db,migrations.FS,".")

	if err!=nil{
		panic(err)
	}

	workoutStore:=store.NewPostgresWorkoutStore(db)


	return &Application{
		Logger: log.New(os.Stdout,"",log.Ldate | log.Ltime),
		WorkoutHandler: api.CreateNewWorkoutHander(workoutStore),
		DB: db,


	}, nil
}


// controller:
func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w,"Health Check\n")
}

