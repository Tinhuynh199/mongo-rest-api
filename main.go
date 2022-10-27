package main

import (
	"fmt"
	"mongorestapi3/internal/app"
	"strconv"

	"context"
)

func main() {
	// Loading Config
	config := app.GetConfig()

	// Loading App
	app := &app.App{}
	app.Initialize(context.Background(), config)

	// Start Server
	server := ""
	if config.Server.Port != nil {
		server = ":" + strconv.FormatInt(*config.Server.Port, 10)
	}
	fmt.Println("Start server")
	app.Run(server)
}
