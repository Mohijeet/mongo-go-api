package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mohijeet/mongo-go-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	oid := bson.ObjectIdHex(id)
	u := models.Users{}
	uc.session.DB("mongo").C("user").FindId(oid).One(&u)
	uj, err := json.Marshal(u)
	fmt.Println(string(uj))
	if err != nil {
		fmt.Print(err)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(uj))
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.Users{}
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Print(u, &u)
	u.Id = bson.NewObjectId()
	uc.session.DB("mongo").C("user").Insert(u)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(uj))

}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	oid := bson.ObjectIdHex(id)
	uc.session.DB("mongo").C("user").RemoveId(oid)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "deleted", string(id))

}
