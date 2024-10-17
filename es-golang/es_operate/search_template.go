package es_operate

import (
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func SearchTemplate(ctx context.Context, name string, client *elastic.Client) error {
	template, err := client.IndexGetIndexTemplate(name).Do(ctx)
	if err != nil {
		return err
	}

	if template == nil {
		return errors.New("template not found")
	}

	for _, indexTemplate := range template.IndexTemplates {
		fmt.Printf("Template Name: %s\n", indexTemplate.Name)
		fmt.Printf("Template Body: %#v\n", indexTemplate.IndexTemplate)
	}

	return nil
}
