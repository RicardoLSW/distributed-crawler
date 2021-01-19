package main

import (
	"distributed-crawler/config"
	"distributed-crawler/engine"
	"distributed-crawler/persist/client"
	"distributed-crawler/scheduler"
	"distributed-crawler/zhenai/parser"
	"fmt"
)

func main() {
	itemSaver, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
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
