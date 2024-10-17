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

func InsertDocWithId(ctx context.Context, id string, name string, doc map[string]interface{}, client *elastic.Client) (*elastic.IndexResponse, error) {
	res, err := client.Index().Index(name).Id(id).BodyJson(doc).Do(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Document Id: %s, Index: %s\n", res.Id, res.Index)
	return res, nil
}
