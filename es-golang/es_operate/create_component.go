package es_operate

import (
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func CreateComponent(ctx context.Context, name string, body map[string]interface{}, client *elastic.Client) error {
	template, err := client.IndexPutComponentTemplate(name).BodyJson(body).Do(ctx)
	if err != nil {
		return err
	}

	if !template.Acknowledged {
		return errors.New("component not acknowledged")
	}

	fmt.Printf("Component %s created\n", name)
	return nil
}
