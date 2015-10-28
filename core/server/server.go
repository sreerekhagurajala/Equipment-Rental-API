package server

import (
	"fmt"
	"strconv"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"github.com/hypebeast/gojistatic"
	"github.com/rs/cors"
	"../routes"
	"../config"
	"../config/database"
	"../router"
)
// Start handles all route configuration and starts the http server
func Start(settings config.Properties, context database.Context) {
	fmt.Println("こんにちは, listening on port :" + strconv.Itoa(settings.Port))

    // Create the main router
	masterRouter := web.New()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
		AllowCredentials: true,
		AllowedHeaders: []string{"*"},
	})

	//Create subroutes
	apiRouter 		:= web.New()
	angularRouter 	:= web.New()

	// Assign sub routes to handle certain path requests
	masterRouter.Handle("/api/*", apiRouter)
	masterRouter.Handle("/*", angularRouter)

	// Apply SubRouter middleware to allow sub routing
	apiRouter.Use(middleware.SubRouter)
	angularRouter.Use(middleware.SubRouter)

	// Serve the static files in the client app directory (this is to host the angular app)
	angularRouter.Use(gojistatic.Static("client/app/", gojistatic.StaticOptions{
		SkipLogging: true,
		Expires: nil,
	}))

	// Apply the CORS options to the main route handler
	masterRouter.Use(c.Handler)

	// Create the routes
	routes.CreateRoutes(router.API{Router:apiRouter, Context:context})

	// Gracefully Serve
	if portIsFree(settings.Port) {
		err := graceful.ListenAndServe(":" + strconv.Itoa(settings.Port), masterRouter)
		if err != nil {
			//		If an error occurs, normally is port is already in use

			//		Don't panic
			panic(err)
		}
	}
}

// Checks if a port is free
func portIsFree(port int) bool {
//	If the port is being used

//	Return false

//	if not in use
	return true
}


