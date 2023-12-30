package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Anand-S23/OneLiner/config"
	"github.com/Anand-S23/OneLiner/internal/controller"
	"github.com/Anand-S23/OneLiner/internal/database"
	"github.com/Anand-S23/OneLiner/internal/router"
	"github.com/Anand-S23/OneLiner/internal/storage"
	"github.com/gorilla/handlers"
)

func main() {
    env, err := config.LoadEnv()
    if err != nil {
        log.Fatal(err)
    }

    db := database.InitDB(10 * time.Second)
    dynamoStore := storage.NewDynamoStore(db)
    controller := controller.NewController(dynamoStore, env.JWT_SECRET, env.PRODUCTION)
    router := router.NewRouter(controller)

	corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

    log.Println("OneLiner running on port: ", env.PORT);
    http.ListenAndServe(":" + env.PORT, corsHandler(router))
}
