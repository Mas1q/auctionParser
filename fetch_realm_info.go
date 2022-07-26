package main

import (
	"context"
	"github.com/hasura/go-graphql-client"
)

type RealmInfo struct {
	Parcels []struct {
		Id                    graphql.String
		EquippedInstallations []struct {
			Id graphql.String
		}
		LastChanneledAlchemica graphql.String
	} `graphql:"parcels(first: 1, where: { id: $id })"`
}

func (*RealmInfo) FetchRealmInfo(id string) (*RealmInfo, error) {
	realm := &RealmInfo{}
	variables := map[string]interface{}{
		"id": graphql.String(id),
	}

	client := graphql.NewClient("https://api.thegraph.com/subgraphs/name/aavegotchi/gotchiverse-matic", nil)

	err := client.Query(context.Background(), &realm, variables)
	if err != nil {
		return realm, err
	}

	return realm, nil
}
