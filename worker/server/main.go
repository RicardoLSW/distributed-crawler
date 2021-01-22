package main

import (
	"distributed-crawler/config"
	"distributed-crawler/rpcsupport"
	"distributed-crawler/worker"
	"fmt"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}))
}
