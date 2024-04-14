package main

import (
	"net/http"
	"time"

	"git.virjar.com/Junhiee/anilismei/initialize"
	"git.virjar.com/Junhiee/anilismei/router"
)

func main() {

	initialize.SetupMysql()

	r := router.Routers()

	s := &http.Server{
		Addr:           ":8081",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
