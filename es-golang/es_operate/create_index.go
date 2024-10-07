package es_operate

import (
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func CreateIndex(ctx context.Context, name string, client *elastic.Client) error {
	// 检查索引是否存在
	exists, err := client.IndexExists(name).Do(ctx)
	if err != nil {
		return err
	}

	if exists {
		errMsg := fmt.Sprintf("Index %s already exists\n", name)
		return errors.New(errMsg)
	}

	// 创建索引
	index, err := client.CreateIndex(name).Do(ctx)
	if err != nil {
		return err
	}

	// 检查索引是否创建成功
	// true表示ES集群的所有节点都接收并处理了创建索引的请求
	// false表示ES集群没有完全确认该操作 可能只有部分节点创建索引成功
	if !index.Acknowledged {
		errMsg := fmt.Sprintf("Index %s creation not acknowledged\n", name)
		return errors.New(errMsg)
	}

	fmt.Printf("Index %s created successfully without predefined mapping\n", name)
	return nil
}

// CreateIndexWithConfig 创建索引并指定索引配置
func CreateIndexWithConfig(ctx context.Context, name string, config string, client *elastic.Client) error {
	exists, err := client.IndexExists(name).Do(ctx)
	if err != nil {
		return err
	}

	if exists {
		errMsg := fmt.Sprintf("Index %s already exists\n", name)
		return errors.New(errMsg)
	}

	index, err := client.CreateIndex(name).BodyString(config).Do(ctx)
	if err != nil {
		return err
	}

	if !index.Acknowledged {
		errMsg := fmt.Sprintf("Index %s creation not acknowledged\n", name)
		return errors.New(errMsg)
	}

	fmt.Printf("Index %s created successfully without predefined mapping\n", name)
	return nil
}
