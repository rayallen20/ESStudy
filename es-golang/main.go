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

	callInsertDocWithId(client)
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

func callCreateTemplate(client *elastic.Client) {
	name := "template_1"
	body := map[string]interface{}{
		"index_patterns": []string{
			"te*",
			"bar*",
		},
		"template": map[string]interface{}{
			"aliases": map[string]interface{}{
				"alias1": struct{}{},
			},

			"settings": map[string]interface{}{
				"number_of_shards": 1,
			},

			"mappings": map[string]interface{}{
				"_source": map[string]interface{}{
					"enabled": false,
				},

				"properties": map[string]interface{}{
					"host_name": map[string]interface{}{
						"type": "keyword",
					},

					"created_at": map[string]interface{}{
						"type":   "date",
						"format": "EEE MMM dd HH:mm:ss Z yyyy",
					},
				},
			},
		},
	}
	err := es_operate.CreateTemplate(context.Background(), name, body, client)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func callMappingCreateComponent(client *elastic.Client) {
	name := "component_mapping_template"
	body := map[string]interface{}{
		"template": map[string]interface{}{
			"mappings": map[string]interface{}{
				"properties": map[string]interface{}{
					"@timestamp": map[string]interface{}{
						"type": "date",
					},

					"host_name": map[string]interface{}{
						"type": "keyword",
					},

					"created_at": map[string]interface{}{
						"type":   "date",
						"format": "EEE MMM dd HH:mm:ss Z yyyy",
					},
				},
			},
		},
	}
	err := es_operate.CreateComponent(context.Background(), name, body, client)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func callSettingCreateComponent(client *elastic.Client) {
	name := "component_setting_template"
	body := map[string]interface{}{
		"template": map[string]interface{}{
			"settings": map[string]interface{}{
				"number_of_shards": 3,
			},

			"aliases": map[string]interface{}{
				"myData": struct{}{},
			},
		},
	}
	err := es_operate.CreateComponent(context.Background(), name, body, client)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func createTemplateBaseOnComponent(client *elastic.Client) {
	name := "my_data_template"
	body := map[string]interface{}{
		"index_patterns": []string{
			"my_data*",
		},

		"priority": 500,

		"composed_of": []string{
			"component_mapping_template",
			"component_setting_template",
		},

		"version": 1,

		"_meta": map[string]interface{}{
			"description": "My custom template",
		},
	}
	err := es_operate.CreateTemplate(context.Background(), name, body, client)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func callSearchTemplate(client *elastic.Client) {
	name := "template_1"
	err := es_operate.SearchTemplate(context.Background(), name, client)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func callDeleteTemplate(client *elastic.Client) {
	name := "template_1"
	err := es_operate.DeleteTemplate(context.Background(), name, client)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func dynamicTemplate(client *elastic.Client) {
	name := "sample_dynamic_template"

	body := map[string]interface{}{
		"index_patterns": []string{
			"sample*",
		},

		"template": map[string]interface{}{
			"mappings": map[string]interface{}{
				"dynamic_templates": []map[string]interface{}{
					{
						"handle_integers": map[string]interface{}{ // handle_integers: 动态模板名称
							"match_mapping_type": "long", // match_mapping_type: 被匹配的、待重新指定的源数据类型
							"mapping": map[string]interface{}{ // mapping: 重新指定的目标数据类型
								"type": "integer",
							},
						},
					},

					{
						"handle_date": map[string]interface{}{
							"match": "date_*", // match: 匹配字段名的通配符
							"mapping": map[string]interface{}{ // mapping: 重新指定的目标数据类型
								"type": "date",
							},
						},
					},
				},
			},
		},
	}

	err := es_operate.CreateTemplate(context.Background(), name, body, client)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}

func callInsertDocWithId(client *elastic.Client) {
	name := "my_index_0501" // 该Index事前在ES中并不存在 但是可以自动创建
	doc := map[string]interface{}{
		"media_array": []string{ // String类型的Array
			"新闻",
			"论坛",
			"博客",
			"电子报",
		},
		"users_array": []struct { // Object类型的Array
			Name string
			Age  int
		}{
			{
				Name: "Mary",
				Age:  12,
			},
			{
				Name: "John",
				Age:  10,
			},
		},
		"size_array": []int{ // long类型的Array
			0,
			50,
			100,
		},
	}
	id := "1"

	_, err := es_operate.InsertDocWithId(context.Background(), id, name, doc, client)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
}
