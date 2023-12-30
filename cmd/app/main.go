package main

import (
	"log"

	"github.com/Anand-S23/OneLiner/config"
)

func main() {
    env, err := config.LoadEnv()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("OneLiner running on port: ", env.PORT);
}
