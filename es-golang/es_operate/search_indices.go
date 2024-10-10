package es_operate

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func SearchMultiIndices(ctx context.Context, client *elastic.Client, indices []string) (*elastic.SearchResult, error) {
	query := elastic.NewMatchAllQuery()

	indicesDoc, err := client.Search().Index(indices...).Query(query).Do(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Found %d documents\n", indicesDoc.TotalHits())

	return indicesDoc, nil
}

func SearchMultiIndicesByExp(ctx context.Context, client *elastic.Client, exp string) (*elastic.SearchResult, error) {
	query := elastic.NewMatchAllQuery()

	indicesDoc, err := client.Search().Index(exp).Query(query).Do(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Found %d documents\n", indicesDoc.TotalHits())

	return indicesDoc, nil
}
