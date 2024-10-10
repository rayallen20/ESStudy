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

func AliasWriteableIndex(ctx context.Context, name string, alias string, client *elastic.Client) error {
	aliasAction := elastic.NewAliasAddAction(alias).Index(name).IsWriteIndex(true)

	aliasIndex, err := client.Alias().Action(aliasAction).Do(ctx)
	if err != nil {
		return err
	}

	if !aliasIndex.Acknowledged {
		msg := fmt.Sprintf("Alias %s was not created for Index %s", alias, name)
		return errors.New(msg)
	}

	fmt.Printf("Alias %s successfully updated. Index %s is now the write index\n", alias, name)
	return nil
}

func GetAliasIndex(ctx context.Context, client *elastic.Client, alias string) (*elastic.AliasesResult, error) {
	aliasesResult, err := client.Aliases().Alias(alias).Do(ctx)
	if err != nil {
		return nil, err
	}

	for indexName, indexResult := range aliasesResult.Indices {
		fmt.Printf("Index: %s\n", indexName)

		// 一个索引可能有多个别名 因此 indexResult.Aliases 是一个slice
		for _, info := range indexResult.Aliases {
			fmt.Printf("Alias name: %s, Is write index: %v\n", info.AliasName, info.IsWriteIndex)
		}
	}

	return aliasesResult, nil
}
