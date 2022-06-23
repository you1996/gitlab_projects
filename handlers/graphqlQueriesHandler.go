package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	graphql "github.com/hasura/go-graphql-client"
	"github.com/you1996/gitlab-projects/backend/model"
)

type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func getProjects(c *gin.Context, tag string) interface{} {
	// initialize global vars
	var cursor string                          // used to track page cursor
	var totalFetchedProjects int               // total of fetched projects
	var starsCount graphql.Int                 // total stars count
	numberOfProjects, err := getQueryString(c) // get the param
	if err != nil {
		return err
	}
	// graphql init
	client := graphql.NewClient("https://gitlab.com/api/graphql", nil)

	// TODO : try to enhance code
	// if we are fetching projects so then tag the function by projects
	if tag == "projects" {
		err := graphqlQueryHelper(c, client, numberOfProjects, cursor, totalFetchedProjects, starsCount, "projects")
		if err != nil {
			return err
		}
	} else {
		// we are fetching stars Count
		err := graphqlQueryHelper(c, client, numberOfProjects, cursor, totalFetchedProjects, starsCount, "stars")
		if err != nil {
			return err
		}
	}
	return nil
}

func graphqlQueryHelper(c *gin.Context, client *graphql.Client, projectsNumber int, cursor string, totalFetchedProjects int, starsCount graphql.Int, tag string) error {
	// variables used to pass dynamic data to graphql query
	variables := map[string]interface{}{
		"numberOfProjects": graphql.Int(projectsNumber),
		"cursor":           graphql.String(cursor),
	}
	// struct to fillin the fetched projects
	var projects struct {
		Projects model.Projects `graphql:"projects(first: $numberOfProjects after:$cursor)"`
	}
	// there is an error from graphql server endpoint
	graphqlError := client.Query(context.Background(), &projects, variables)
	if graphqlError != nil {
		return graphqlError
	}
	// case fetching projects
	if tag == "projects" {
		// create a map so can be like the json task
		projectsMap := make(map[graphql.String]graphql.Int, projectsNumber)
		for _, value := range projects.Projects.Nodes {
			// key:value (name:stars)
			projectsMap[value.Name] = value.StarCount
			totalFetchedProjects++
		}
		// TODO : for frontend
		dat, err := json.Marshal(projectsMap)
		if err != nil {
			panic(err)
		}
		// case when we should fetch more pages
		if projects.Projects.PageInfo.HasNextPage == true && totalFetchedProjects < projectsNumber {
			// server send event to client side
			c.SSEvent("", gin.H{
				"data": string(dat),
			})
			// recursion but changing the cursor
			graphqlQueryHelper(c, client, projectsNumber, string(projects.Projects.PageInfo.EndCursor), totalFetchedProjects, starsCount, tag)
		} else {
			// Nomore data to fetch
			c.SSEvent("", gin.H{
				"data": string(dat),
			})
			return nil
		}
	} else {
		// case fetching stars count only
		for _, value := range projects.Projects.Nodes {
			// add stars to the global starsCOunt variable
			starsCount += value.StarCount
			totalFetchedProjects++
		}
		if projects.Projects.PageInfo.HasNextPage == true && totalFetchedProjects < projectsNumber {
			c.SSEvent("message", starsCount)
			graphqlQueryHelper(c, client, projectsNumber, string(projects.Projects.PageInfo.EndCursor), totalFetchedProjects, starsCount, tag)
		} else {
			c.SSEvent("message", starsCount)
			return nil
		}
	}
	return nil
}

func getQueryString(c *gin.Context) (int, interface{}) {
	var parsedError *apiError
	// get the param
	queryParamAsString := c.Query("number-of-projects")
	// check if it's not empty
	if queryParamAsString == "" {

		parsedError = &apiError{Code: http.StatusInternalServerError, Message: "Invalid parameter !"}
		return 0, parsedError
	}
	// convert to integer
	queryParamAsInt, err := strconv.Atoi(queryParamAsString)
	// check if it's not a lettres string
	if err != nil {
		parsedError = &apiError{Code: http.StatusInternalServerError, Message: "Can't parse query string !"}
		return 0, parsedError
	}
	// check if it's not zero
	if queryParamAsInt == 0 {
		parsedError = &apiError{Code: http.StatusInternalServerError, Message: "Number of projects to fetch cannot be zero !"}
		return 0, parsedError
	}
	return queryParamAsInt, nil
}
