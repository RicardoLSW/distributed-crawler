package main

import (
	"distributed-crawler/config"
	"distributed-crawler/engine"
	itemSaver "distributed-crawler/persist/client"
	"distributed-crawler/scheduler"
	worker "distributed-crawler/worker/client"
	"distributed-crawler/zhenai/parser"
	"fmt"
)

func main() {
	itemChan, err := itemSaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	//concurrentEngine.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	concurrentEngine.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun/shanghai",
		Parser: engine.NewFuncParser(parser.ParseCity, config.ParseCity),
	})
}
