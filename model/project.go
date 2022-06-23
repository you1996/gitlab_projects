package model

import (
	graphql "github.com/hasura/go-graphql-client"
)

// one project
type Project struct {
	Name      graphql.String `json:"name"`
	StarCount graphql.Int    `json:"star"`
}