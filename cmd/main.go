package main

import (
	"github.com/gin-gonic/gin"

	"github.com/you1996/gitlab-projects/backend/handlers"
)


func setupRouter() *gin.Engine {
	r := gin.Default()
	// Creating a route that is going to be called when we make a GET request to the endpoint
	// `/projects-names-and-stars`
	r.GET("/projects-names-and-stars", handlers.GetLastNProjectsWithStarsAndNames)

	// A route that is going to be called when we make a GET request to the endpoint
	// `/projects-total-stars`
	r.GET("/projects-total-stars", handlers.GetSumOfStarsForLastNProjects)

	return r
}

// It creates a new router, sets up the routes, and then starts the server
func main() {
	r := setupRouter()
	r.Run(":8082")
}
