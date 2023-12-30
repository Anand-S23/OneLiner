package main

import (
	"context"
	"log"
	"time"

	"github.com/Anand-S23/OneLiner/config"
	"github.com/Anand-S23/OneLiner/internal/database"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
    env, err := config.LoadEnv()
    if err != nil {
        log.Fatal(err)
    }

    db := database.InitDB(10 * time.Second)
    _, err = db.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
    if err != nil {
        log.Fatal(err)
    }

    log.Println("OneLiner running on port: ", env.PORT);
}
