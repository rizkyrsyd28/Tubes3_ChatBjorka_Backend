package main

import (
	"fmt"
	"gopkg.in/fsnotify.v1"
	"os"
	"strings"

	"github.com/rizkyrsyd28/internal/route"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()

	// Add CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "https://saddamannais.github.io/Tubes3_13521109/"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))

	route.Routes(r)

	// creates a new file watcher for App_offline.htm
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	// watch for App_offline.htm and exit the program if present
	// This allows continuous deployment on App Service as the .exe will not be
	// terminated otherwise
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if strings.HasSuffix(event.Name, "App_offline.htm") {
					fmt.Println("Exiting due to App_offline.htm being present")
					os.Exit(0)
				}
			}
		}
	}()

	// get the current working directory and watch it
	currentDir, err := os.Getwd()
	if err := watcher.Add(currentDir); err != nil {
		fmt.Println("ERROR", err)
	}

	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	r.Run(":" + port)
}
