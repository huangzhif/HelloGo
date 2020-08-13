package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetJsonContent(jfile string) []uint8 {
	/*
		获取json文件内容
	*/
	jsonFile, err := os.Open(jfile)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue

}

func PathExist(path string) bool {
	/*
		判断文件夹或文件是否存在
	*/
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// func main() {
// 	ret := GetJsonContent("test.json")

// 	var tconfigs Tcconfigs

// 	json.Unmarshal(ret, &tconfigs)

// 	for i := 0; i < len(tconfigs.Tcconfigs); i++ {
// 		fmt.Println("AK:" + tconfigs.Tcconfigs[i].AK)
// 		fmt.Println("SK:" + tconfigs.Tcconfigs[i].SK)
// 		fmt.Println("Mainpart:" + tconfigs.Tcconfigs[i].Mainpart)
// 	}
// }

