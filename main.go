package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github.com/brandonspitz/MongoDB-API/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:00001", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:00000")
	if err != nil {
		panic(err)
	}
	return s
}
