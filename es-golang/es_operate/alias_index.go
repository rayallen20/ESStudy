package es_operate

import (
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func AliasIndex(ctx context.Context, name string, alias string, client *elastic.Client) error {
	// ES中,1个索引可以有多个别名,1个别名也可以对应多个索引
	// 因此,在为索引指定别名前,不需要检查别名是否存在
	// 但是,在删除别名时,需要检查别名是否存在

	// 为索引指定别名
	aliasIndex, err := client.Alias().Add(name, alias).Do(ctx)
	if err != nil {
		return err
	}

	if !aliasIndex.Acknowledged {
		msg := fmt.Sprintf("Alias %s was not created for Index %s", alias, name)
		return errors.New(msg)
	}

	fmt.Printf("Alias %s was created for Index %s\n", alias, name)
	return nil
}
