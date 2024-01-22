package router

import (
	"net/http"

	"github.com/Anand-S23/Snippet/internal/controller"
	"github.com/Anand-S23/Snippet/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(c *controller.Controller) *mux.Router {
    router := mux.NewRouter()
	router.HandleFunc("/ping", HandleFunc(c.Ping))

    // Auth
    router.HandleFunc("/register", HandleFunc(c.SignUp)).Methods("POST")
    router.HandleFunc("/login", HandleFunc(c.Login)).Methods("POST")
    router.HandleFunc("/logout", HandleFunc(c.Logout)).Methods("POST")

    // Post
    router.HandleFunc("/upload", middleware.Authentication(HandleFunc(c.UploadFile), c.JwtSecretKey)).Methods("POST")
    router.HandleFunc("/post/create", middleware.Authentication(HandleFunc(c.CreatePost), c.JwtSecretKey)).Methods("POST")
    router.HandleFunc("/post/get/{id}", HandleFunc(c.ReadPost)).Methods("GET")
    router.HandleFunc("/post/delete/{id}", middleware.Authentication(HandleFunc(c.DeletePost), c.JwtSecretKey)).Methods("POST")

    return router
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func HandleFunc(fn apiFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := fn(w, r)
        if err != nil {
            errMsg := map[string]string { "error": err.Error() }
            controller.WriteJSON(w, http.StatusInternalServerError, errMsg)
        }
    }
}

