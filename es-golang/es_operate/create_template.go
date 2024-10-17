package es_operate

import (
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func CreateTemplate(ctx context.Context, name string, body map[string]interface{}, client *elastic.Client) error {
	template, err := client.IndexPutIndexTemplate(name).BodyJson(body).Do(ctx)
	if err != nil {
		return err
	}

	if !template.Acknowledged {
		return errors.New("template not acknowledged")
	}

	fmt.Printf("Index template %s created\n", name)
	return nil
}
