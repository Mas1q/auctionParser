package main

import (
	"context"
	"github.com/hasura/go-graphql-client"
)

var realm struct {
	Parcels []struct {
		Id                    graphql.String
		EquippedInstallations []struct {
			Id graphql.String
		}
		LastChanneledAlchemica graphql.String
	} `graphql:"parcels(first: 1, where: { id: $id })"`
}

func FetchRealm(id string) interface{} {
	variables := map[string]interface{}{
		"id": graphql.String(id),
	}

	client := graphql.NewClient("https://api.thegraph.com/subgraphs/name/aavegotchi/gotchiverse-matic", nil)
	err := client.Query(context.Background(), &realm, variables)
	if err != nil {
		panic(err)
	}

	return realm
}
