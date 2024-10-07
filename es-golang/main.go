package main

import (
	"context"
	"es-go/es_operate"
	"fmt"
)

func main() {
	address := "http://192.168.1.195:9200"
	client, err := es_operate.Conn(address)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	//indexName := "my_test_index_1"
	//err = es_operate.CreateIndex(context.Background(), indexName, client)
	//if err != nil {
	//	fmt.Printf("%s\n", err.Error())
	//}

	indexName := "hamlet-1"
	//	config := `
	//{
	//    "settings": {
	//        "number_of_shards": 2,
	//        "number_of_replicas": 1
	//    },
	//    "mappings": {
	//        "properties": {
	//            "cont": {
	//                "type": "text",
	//                "analyzer": "standard",
	//                "fields": {
	//                    "field": {
	//                        "type": "keyword"
	//                    }
	//                }
	//            }
	//        }
	//    },
	//    "aliases": {
	//        "hamlet": {}
	//    }
	//}
	//`
	//	err = es_operate.CreateIndexWithConfig(context.Background(), indexName, config, client)
	//	if err != nil {
	//		fmt.Printf("%s\n", err.Error())
	//	}

	updateSettings := `
{
    "index": {
        "number_of_replicas": 2,
        "refresh_interval": "30s"
    }
}
`
	err = es_operate.UpdateSettings(context.Background(), indexName, updateSettings, client)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
