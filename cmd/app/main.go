package main

import (
	"log"
	"time"

	"github.com/Anand-S23/OneLiner/config"
	"github.com/Anand-S23/OneLiner/internal/controller"
	"github.com/Anand-S23/OneLiner/internal/database"
	"github.com/Anand-S23/OneLiner/internal/router"
	"github.com/Anand-S23/OneLiner/internal/storage"
)

func main() {
    env, err := config.LoadEnv()
    if err != nil {
        log.Fatal(err)
    }

    db := database.InitDB(10 * time.Second)
    dynamoStore := storage.NewDynamoStore(db)
    controller := controller.NewController(dynamoStore, "", env.PRODUCTION)
    router := router.NewRouter(controller)

    log.Println("OneLiner running on port: ", env.PORT);
}
