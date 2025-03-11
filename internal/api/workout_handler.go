package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/tanvir-rifat007/workoutApi/internal/store"
)


type WorkoutHandler struct{
	WorkoutStore store.WorkoutStore
}


func CreateNewWorkoutHander(workoutStore store.WorkoutStore) *WorkoutHandler{
	return &WorkoutHandler{
		WorkoutStore: workoutStore,
	}
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

func (wh *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request){
  var workout store.Workout

	err:=json.NewDecoder(r.Body).Decode(&workout)

	if err!=nil{
		fmt.Println(err)
		http.Error(w,"Bad Request",http.StatusBadRequest)
		return
	}

	createdWorkout,err:=wh.WorkoutStore.CreateWorkout(&workout)

	if err!=nil{
		fmt.Println(err)
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(createdWorkout)

}


