package main

import (
	"context"
	"es-go/es_operate"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func main() {
	address := "http://192.168.1.195:9200"
	client, err := es_operate.Conn(address)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	callGetAliasIndex(client)
}

func callCreateIndex(client *elastic.Client) {
	indexName := "my_logs_202410"
	err := es_operate.CreateIndex(context.Background(), indexName, client)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func callCreateIndexWithConfig(client *elastic.Client) {
	indexName := "hamlet-1"
	config := `
	{
	   "settings": {
	       "number_of_shards": 2,
	       "number_of_replicas": 1
	   },
	   "mappings": {
	       "properties": {
	           "cont": {
	               "type": "text",
	               "analyzer": "standard",
	               "fields": {
	                   "field": {
	                       "type": "keyword"
	                   }
	               }
	           }
	       }
	   },
	   "aliases": {
	       "hamlet": {}
	   }
	}
	`
	err := es_operate.CreateIndexWithConfig(context.Background(), indexName, config, client)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func callUpdateSettings(client *elastic.Client) {
	indexName := "hamlet-1"

	updateSettings := `
{
    "index": {
        "number_of_replicas": 2,
        "refresh_interval": "30s"
    }
}
`
	err := es_operate.UpdateSettings(context.Background(), indexName, updateSettings, client)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func callDeleteIndex(client *elastic.Client) {
	indexName := "my_test_index_1"
	err := es_operate.DeleteIndex(context.Background(), indexName, client)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func callCleanIndex(client *elastic.Client) {
	indexName := "hamlet-1"
	err := es_operate.CleanIndex(context.Background(), indexName, client)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func callAliasIndex(client *elastic.Client) {
	indexName := "my_logs_202410"
	aliasName := "my_logs"
	err := es_operate.AliasIndex(context.Background(), indexName, aliasName, client)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func callGetIndexInfo(client *elastic.Client) {
	indexName := "hamlet-1"
	indexInfo, err := es_operate.GetIndexInfo(context.Background(), indexName, client)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	for index, info := range indexInfo {
		fmt.Printf("Index: %s\n", index)
		fmt.Printf("Aliases: %v\n", info.Aliases)
		fmt.Printf("Mappings: %v\n", info.Mappings)
		fmt.Printf("Settings: %v\n", info.Settings)
	}
}

func callSearchMultiIndices(client *elastic.Client) {
	indices := []string{"my_logs_202409", "my_logs_202410"}
	indicesDoc, err := es_operate.SearchMultiIndices(context.Background(), client, indices)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	for _, hit := range indicesDoc.Hits.Hits {
		fmt.Printf("DocumentId: %s, Source: %s\n", hit.Id, hit.Source)
	}
}

func callSearchMultiIndicesByExp(client *elastic.Client) {
	indicesExp := "my_logs*"
	indicesDoc, err := es_operate.SearchMultiIndicesByExp(context.Background(), client, indicesExp)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	for _, hit := range indicesDoc.Hits.Hits {
		fmt.Printf("DocumentId: %s, Source: %s\n", hit.Id, hit.Source)
	}
}

func callSearchIndex(client *elastic.Client) {
	name := "my_logs"
	indicesDoc, err := es_operate.SearchIndex(context.Background(), client, name)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	for _, hit := range indicesDoc.Hits.Hits {
		fmt.Printf("DocumentId: %s, Source: %s\n", hit.Id, hit.Source)
	}
}

func callInsertDoc(client *elastic.Client) {
	name := "my_logs"
	doc := map[string]interface{}{
		"index": struct{}{},
		"title": "001",
	}

	_, err := es_operate.InsertDoc(context.Background(), client, name, doc)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
}

func callAliasWriteableIndex(client *elastic.Client) {
	name := "my_logs_202409"
	alias := "my_logs"
	err := es_operate.AliasWriteableIndex(context.Background(), name, alias, client)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func callGetAliasIndex(client *elastic.Client) {
	alias := "my_logs"
	_, err := es_operate.GetAliasIndex(context.Background(), client, alias)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
