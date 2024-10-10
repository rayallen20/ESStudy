package es_operate

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func InsertDoc(ctx context.Context, client *elastic.Client, name string, doc map[string]interface{}) (*elastic.IndexResponse, error) {
	res, err := client.Index().Index(name).BodyJson(doc).Do(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Document Id: %s, Index: %s\n", res.Id, res.Index)
	return res, nil
}
