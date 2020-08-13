package main

import (
	"encoding/json"
	"fmt"
	"HelloGo/utils"

	"github.com/robfig/cron"

	"HelloGo/tencent"
)

func test() {
	fmt.Println("test:")

	ret := utils.GetJsonContent("")
	var tconfigs utils.Tcconfigs

	json.Unmarshal(ret, &tconfigs)

	for i := 0; i < len(tconfigs.Tcconfigs); i++ {
		fmt.Println("AK:" + tconfigs.Tcconfigs[i].AK)
		fmt.Println("SK:" + tconfigs.Tcconfigs[i].SK)
		fmt.Println("Mainpart:" + tconfigs.Tcconfigs[i].Mainpart)
	}
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
