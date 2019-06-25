package app

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"golang_rabbitmq/rest_api_server/app/handler"
	"golang_rabbitmq/model"
	"golang_rabbitmq/config"
	"golang_rabbitmq/rest_api_server/app/message_handler"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	//dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
	//	config.DB.Username,
	//	config.DB.Password,
	//	config.DB.Name,
	//	config.DB.Charset)

	//db, err := gorm.Open(config.DB.Dialect, dbURI)
	db, err := gorm.Open(config.DB.Dialect, config.DB.Name)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	message_handler.Initialize()
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {


	// Routing for handling the tasks
	a.Get("/tasks", a.GetAllTasks)
	a.Post("/tasks", a.CreateTask)
	a.Get("/tasks/{id:[0-9]+}", a.GetTask)
	a.Put("/tasks/{id:[0-9]+}", a.UpdateTask)
	a.Delete("/tasks/{id:[0-9]+}", a.DeleteTask)
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

/*
** Tasks Handlers
 */
func (a *App) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	handler.GetAllTasks(a.DB, w, r)
}

func (a *App) CreateTask(w http.ResponseWriter, r *http.Request) {
	handler.CreateTask(a.DB, w, r)
}

func (a *App) GetTask(w http.ResponseWriter, r *http.Request) {
	handler.GetTask(a.DB, w, r)
}

func (a *App) UpdateTask(w http.ResponseWriter, r *http.Request) {
	handler.UpdateTask(a.DB, w, r)
}

func (a *App) DeleteTask(w http.ResponseWriter, r *http.Request) {
	handler.DeleteTask(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
