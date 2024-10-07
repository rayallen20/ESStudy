package es_operate

import (
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func DeleteIndex(ctx context.Context, name string, client *elastic.Client) error {
	deletedIndex, err := client.DeleteIndex(name).Do(ctx)
	if err != nil {
		return err
	}

	if !deletedIndex.Acknowledged {
		msg := fmt.Sprintf("Index %s was not deleted", name)
		return errors.New(msg)
	}

	fmt.Printf("Index %s was deleted\n", name)
	return nil
}

func CleanIndex(ctx context.Context, name string, client *elastic.Client) error {
	// 构建删除所有文档的查询
	query := elastic.NewMatchAllQuery()

	// 执行 _delete_by_query API
	deleteByQuery, err := client.DeleteByQuery().Index(name).Query(query).Do(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("Deleted %d documents from Index %s\n", deleteByQuery.Deleted, name)
	return nil
}
