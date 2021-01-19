package main

import (
	"distributed-crawler/engine"
	"distributed-crawler/persist/client"
	"distributed-crawler/scheduler"
	"distributed-crawler/zhenai/parser"
)

func main() {
	itemSaver, err := client.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}
	concurrentEngine := engine.ConcurrentEngine{Scheduler: &scheduler.QueuedScheduler{}, WorkerCount: 10, ItemChan: itemSaver}
	//concurrentEngine.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	concurrentEngine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
