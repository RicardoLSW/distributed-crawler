package persist

import (
	"context"
	"distributed-crawler/engine"
	"errors"

	"github.com/olivere/elastic/v7"
)

func Save(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.
		Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
