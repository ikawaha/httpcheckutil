package main

import (
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/ikawaha/gqlgen-todos/graph"
	"github.com/ikawaha/httpcheckutil/graphql"
)

func TestServer(t *testing.T) {
	h := handler.New(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)
	h.AddTransport(transport.POST{})

	checker := graphql.New(h, graphql.Debug())
	checker.Test(t).
		WithHeader("Content-Type", "application/json").
		Query(`query {todos {text}}`).
		Check().
		HasStatusOK().
		HasNoErrors().
		HasData(map[string]any{
			"todos": []any{},
		})
}
