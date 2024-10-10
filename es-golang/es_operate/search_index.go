package es_operate

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func SearchIndex(ctx context.Context, client *elastic.Client, name string) (*elastic.SearchResult, error) {
	query := elastic.NewMatchAllQuery()

	indicesDoc, err := client.Search().Index(name).Query(query).Do(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Found %d documents\n", indicesDoc.TotalHits())

	return indicesDoc, nil
}
