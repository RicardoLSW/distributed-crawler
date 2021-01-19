package main

import (
	"distributed-crawler/persist"
	"distributed-crawler/rpcsupport"
	"log"

	"github.com/olivere/elastic/v7"
)

func main() {
	log.Fatal(serveRpc(":1234", "dating_profile"))
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
