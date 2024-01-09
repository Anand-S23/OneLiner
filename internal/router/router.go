package router

import (
	"net/http"

	"github.com/Anand-S23/Snippet/internal/controller"
	"github.com/gorilla/mux"
)

func NewRouter(c *controller.Controller) *mux.Router {
    router := mux.NewRouter()
	router.HandleFunc("/ping", HandleFunc(c.Ping))

    router.HandleFunc("/register", HandleFunc(c.SignUp)).Methods("POST")
    router.HandleFunc("/login", HandleFunc(c.Login)).Methods("POST")
    router.HandleFunc("/logout", HandleFunc(c.Logout)).Methods("POST")

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

