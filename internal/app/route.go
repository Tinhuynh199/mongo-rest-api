package app

import (
	"log"
	"net/http"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

// // Set all required routers
func (a *App) setRouters() {

	// Routing for handling the projects
	userPath := "/students"
	a.Get(userPath, a.StudentHandler.GetAllStudents)
	a.Get(userPath+"/{id}", a.StudentHandler.GetStudent)
	a.Post(userPath, a.StudentHandler.InsertStudent)
	a.Put(userPath+"/{id}", a.StudentHandler.UpdateStudent)
	a.Delete(userPath+"/{id}", a.StudentHandler.DeleteStudent)
}

// // Wrap  the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.StrictSlash(true).HandleFunc(path, f).Methods(GET)
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.StrictSlash(false).HandleFunc(path, f).Methods(POST)
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.StrictSlash(false).HandleFunc(path, f).Methods(PUT)
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.StrictSlash(false).HandleFunc(path, f).Methods(DELETE)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
