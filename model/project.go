package model

import (
	graphql "github.com/hasura/go-graphql-client"
)

// many projects
type Projects struct {
	Count    graphql.Int `json:"count"`
	Nodes    []Project   `json:"node"`
	PageInfo PageInfo    `json:"pageinf"`
}

// one project
type Project struct {
	Name      graphql.String `json:"name"`
	StarCount graphql.Int    `json:"star"`
}

// page info type
type PageInfo struct {
	EndCursor   graphql.String
	HasNextPage graphql.Boolean
}
