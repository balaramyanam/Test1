package controllers
 
import (
	"encoding/json"

	"net/http"
	"instance/model"

	"github.com/gorilla/mux"
)

type Medicine interface {
	Search(res http.ResponseWriter, req *http.Request)
	Signup(res http.ResponseWriter, req *http.Request)
	Order(res http.ResponseWriter, req *http.Request)
}

type Api struct {
	ttd model.DBV // interface objects

	// pointer to struct
}

func Createinstance() *Api {

	return &Api{}

}

func Signup(res http.ResponseWriter, req *http.Request) {

}

func Order(res http.ResponseWriter, req *http.Request) {

}

func Search(res http.ResponseWriter, req *http.Request) {

	ds := mux.Vars(req)

	copy := ds["name"]

	if copy != "" {

	}

	pradeep := model.Idlibatch{1, 25, "coke"}

	//p.ttd.Save(&pradeep)

	json.NewEncoder(res).Encode(pradeep)

}
