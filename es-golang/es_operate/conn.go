package es_operate

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func Conn(address string) (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(address),
		elastic.SetSniff(false),
	)

	if err != nil {
		return nil, err
	}

	err = Ping(context.Background(), client, address)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func Ping(ctx context.Context, client *elastic.Client, address string) error {
	// 使用Ping命令检测ES集群是否连接成功
	info, code, err := client.Ping(address).Do(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("ES return code %d and version %s\n", code, info.Version.Number)

	// 健康检查
	health, err := client.ClusterHealth().Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Cluster health Status: %s\n", health.Status)

	return nil
}
