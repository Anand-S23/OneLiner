package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Anand-S23/Snippet/config"
	"github.com/Anand-S23/Snippet/internal/blob"
	"github.com/Anand-S23/Snippet/internal/controller"
	"github.com/Anand-S23/Snippet/internal/database"
	"github.com/Anand-S23/Snippet/internal/router"
	"github.com/Anand-S23/Snippet/internal/storage"
	"github.com/gorilla/handlers"
)

func main() {
    env, err := config.LoadEnv()
    if err != nil {
        log.Fatal(err)
    }

    timeout := 10 * time.Second
    db := database.InitDB(timeout)
    s3 := blob.InitBlob(env.S3_BUCKET, timeout)
    store := storage.NewSnippetStore(db, database.SnippetTableName, s3)
    controller := controller.NewController(store, env.JWT_SECRET, env.PRODUCTION)
    router := router.NewRouter(controller)

	corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

    log.Println("Snippet running on port: ", env.PORT);
    http.ListenAndServe(":" + env.PORT, corsHandler(router))
}

