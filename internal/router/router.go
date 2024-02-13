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
    router.HandleFunc("/getAuthUserID", middleware.Authentication(HandleFunc(c.GetAuthUserID), c.JwtSecretKey, c.CookieSecret)).Methods("GET")

    // S3
    router.HandleFunc("/upload", middleware.Authentication(HandleFunc(c.UploadFiles), c.JwtSecretKey, c.CookieSecret)).Methods("POST")
    router.HandleFunc("/getFiles", HandleFunc(c.GetFiles)).Methods("POST")

    // Post
    router.HandleFunc("/posts", middleware.Authentication(HandleFunc(c.GetPostsForCurrentUser), c.JwtSecretKey, c.CookieSecret)).Methods("GET")
    router.HandleFunc("/post/create", middleware.Authentication(HandleFunc(c.CreatePost), c.JwtSecretKey, c.CookieSecret)).Methods("POST")
    router.HandleFunc("/post/read/{id}", HandleFunc(c.ReadPost)).Methods("GET")
    router.HandleFunc("/post/update/{id}", middleware.Authentication(HandleFunc(c.UpdatePost), c.JwtSecretKey, c.CookieSecret)).Methods("POST")
    router.HandleFunc("/post/delete/{id}", middleware.Authentication(HandleFunc(c.DeletePost), c.JwtSecretKey, c.CookieSecret)).Methods("POST")

    return router
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func HandleFunc(fn apiFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := fn(w, r)
        if err != nil {
            errMsg := controller.ErrorMessage(err.Error())
            controller.WriteJSON(w, http.StatusInternalServerError, errMsg)
        }
    }
}

