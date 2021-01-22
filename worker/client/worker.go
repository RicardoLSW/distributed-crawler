package client

import (
	"distributed-crawler/config"
	"distributed-crawler/engine"
	"distributed-crawler/rpcsupport"
	"distributed-crawler/worker"
	"fmt"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}

	return func(r engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(r)
		var sResult worker.ParseResult
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}, nil
}
