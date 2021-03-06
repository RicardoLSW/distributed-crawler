package main

import (
	"distributed-crawler/config"
	"distributed-crawler/persist"
	"distributed-crawler/rpcsupport"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaveService{
		Client: client,
		Index:  index,
	})
}
