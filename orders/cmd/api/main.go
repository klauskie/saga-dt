package main

import (
	"github.com/klauskie/saga-dt/orders/config"
	"github.com/klauskie/saga-dt/orders/server"
	"log"
)

func main() {
	env := config.DefaultEnv()
	app := server.NewApp(env)

	if err := app.Run("8089"); err != nil {
		log.Fatalf("server could not initialize")
	}
}
