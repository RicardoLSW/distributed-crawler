package worker

import (
	"distributed-crawler/engine"
	"fmt"
)

type CrawlService struct {
}

func (CrawlService) Process(req Request, result *ParseResult) error {
	fmt.Println(req)
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}

	*result = SerializeResult(engineResult)
	return nil
}
