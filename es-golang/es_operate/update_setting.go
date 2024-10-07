package es_operate

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func UpdateSettings(ctx context.Context, name string, settings string, client *elastic.Client) error {
	_, err := client.IndexPutSettings(name).BodyString(settings).Do(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("Index %s settings updated successfully\n", name)
	return nil
}
