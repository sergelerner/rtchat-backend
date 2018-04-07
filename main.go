package main

import (
	"log"
	"net/http"

	r "gopkg.in/gorethink/gorethink.v4"
)

type Channel struct {
	Id   string `json:"id" gorethink:"id,omitempty"`
	Name string `json:"name" gorethink:"name"`
}

type User struct {
	Id   string `gorethink: "id,omitempty"`
	Name string `gorethink: "name"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "rtc",
	})

	if err != nil {
		log.Panic(err.Error())
	}

	router := NewRouter(session)
	router.Handle("channel add", AddChannel)
	router.Handle("channel subscribe", SubscribeChannel)
	router.Handle("channel unsubscribe", UnsubscribeChannel)
	http.Handle("/", router)
	http.ListenAndServe(":4000", nil)
}
