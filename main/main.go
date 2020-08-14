package main

import (
	"fmt"
	"os"

	"github.com/robfig/cron"

	"HelloGo/tencent"
)

func test() {
    fmt.Println(os.Getenv("PROJPATH"))
}

func main() {
	c := cron.New()
	c.AddFunc("0 */60 9-19 * * *", tencent.Calltcapi)

	go c.Start()
	// defer c.Stop()

	//阻塞，让main函数不退出，在后台一直执行
	select {
	// case <-time.After(time.Second * 10):
	// 	return
	}
}
