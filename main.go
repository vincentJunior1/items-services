package main

import (
	"fmt"
	"items/helpers"
	"items/routes"
	"os"
	"os/signal"
	"syscall"

	ctrl "items/controllers"
	"items/model"

	"github.com/joho/godotenv"
	//  nama module kita/nama file yang kita mau import
)

func main() {
	fmt.Println("Hello world")
	godotenv.Load(".env")
	// Depedency Injection
	log := helpers.InitializeLogging()
	model := model.InitDatabase()
	ctrls := ctrl.InitControllers(model)
	route := routes.InitRoutes(ctrls, log)
	serverErr := make(chan error, 1)

	go func() {
		serverErr <- route.StartServer()
	}()

	var signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	if err := <-signalChan; err != nil {
		log.Println("Shutdown server, caused by:", err)
	}
}

// View (Routing api) -> isinya route
// Controllers (business logic) ->
// Model (Model database kita)

// Routes -> Controllers -> Model
// Models -> Controllers -> Response
