package tencent

import (
	"encoding/json"
	"fmt"
	"HelloGo/utils"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func Calltcapi() {
	f, err := os.OpenFile("logrus.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//完成后，延迟关闭
	defer f.Close()
	// 设置日志输出到文件
	log.SetOutput(f)
	log.Println("Calltcapi")

	YEAR := time.Now().Year()
	MONTH := time.Now().Month()
	DAY := time.Now().Day()

	var thisMonth int
	var wg sync.WaitGroup

	//获取项目目录
	projdir := os.Getenv("PROJPATH")

	if !utils.PathExist(projdir + "/billinfo/files_dir/tencent") {
		os.MkdirAll(projdir+"/billinfo/files_dir/tencent", 0755)
	}

	//读取ak/sk配置文件，格式：
	/*
			{
			"TCconfigs": [{
				"AK": "ak",
				"SK": "sk",
				"mainpart": "主体"
			}]
		}
	*/
	ret := utils.GetJsonContent("../tencent/conf.json")
	var tconfigs utils.Tcconfigs

	json.Unmarshal(ret, &tconfigs)
	for cf := 0; cf < len(tconfigs.Tcconfigs); cf++ {
		for year := 2020; year <= YEAR; year++ {
			if year == YEAR {
				thisMonth = int(MONTH)
			} else {
				thisMonth = 12
			}

			for month := 1; month <= thisMonth; month++ {
				bc := strconv.Itoa(year) + "-" + fmt.Sprintf("%02d", month)
				filename := tconfigs.Tcconfigs[cf].Mainpart + "-" + bc + ".json"
				abpath := projdir+"/billinfo/files_dir/tencent/"+ filename

				// 云厂商api限制，防止一秒内太频繁地调用
				time.Sleep(1 * time.Second)

				// 条件：文件不存在 || 当月前五天更新上个月 || 当月 || 跨年 当年一月份前五天更新去年12月份的
				if !utils.PathExist(abpath) || (year == YEAR && month == int(MONTH)-1 && DAY <= 5) || (year == YEAR && month == int(MONTH)) || (year == YEAR-1 && month == 12 && DAY <= 5 && int(MONTH) == 1) {
					wg.Add(1)

					go func(ak, sk, bc, abpath string) {
						defer wg.Done()
						callsdk(ak, sk, bc, abpath)
					}(tconfigs.Tcconfigs[cf].AK, tconfigs.Tcconfigs[cf].SK, bc, abpath)

				}
			}
		}
	}

	wg.Wait()
}

func callsdk(ak, sk, bc, filepath string) {

	credential := common.NewCredential(
		ak,
		sk)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "billing.tencentcloudapi.com"
	client, _ := billing.NewClient(credential, "", cpf)

	request := billing.NewDescribeBillResourceSummaryRequest()

	params := "{\"Offset\":0,\"Limit\":1000,\"PeriodType\":\"byPayTime\",\"Month\":\"" + bc + "\"}"
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeBillResourceSummary(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	ret := response.ToJsonString()

	fp, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()
	_, err = fp.WriteString(ret)

	if err != nil {
		fmt.Println(err)
	}

}

//func main() {
//	Calltcapi()
//}

