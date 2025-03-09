package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)


type WorkoutHandler struct{}


func CreateNewWorkoutHander() *WorkoutHandler{
	return &WorkoutHandler{}
}


func (workout *WorkoutHandler) GetWorkoutById(w http.ResponseWriter, r *http.Request){
	 id:= chi.URLParam(r,"id")

	if id==""{
		http.NotFound(w,r)
		return
	}

	// convert the string id to int:

	workoutId,err:= strconv.ParseInt(id,10,64)

	if err!=nil{
		http.NotFound(w,r)
		return 
	}

	fmt.Fprintf(w,"The id is : %d\n",workoutId)


}

func (workout *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Workout is created")
	
}


