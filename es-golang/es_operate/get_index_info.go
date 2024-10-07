package es_operate

import (
	"context"
	"github.com/olivere/elastic/v7"
)

func GetIndexInfo(ctx context.Context, name string, client *elastic.Client) (map[string]*elastic.IndicesGetResponse, error) {
	return client.IndexGet(name).Do(ctx)
}
