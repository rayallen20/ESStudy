package es_operate

import (
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func DeleteTemplate(ctx context.Context, name string, client *elastic.Client) error {
	deletedTemplate, err := client.IndexDeleteIndexTemplate(name).Do(ctx)
	if err != nil {
		return err
	}

	if !deletedTemplate.Acknowledged {
		return errors.New("delete template not acknowledged")
	}

	fmt.Printf("Template %s deleted\n", name)

	return nil
}
