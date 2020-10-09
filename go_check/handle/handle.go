package handle

import (
	"JccApiTest/go_check/common"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
)

func HandleReq(req *httplib.BeegoHTTPRequest) map[string]interface{} {
	str, err := req.String()
	if err != nil {
		panic(err)
	}

	var outPutData map[string]interface{}
	if err := json.Unmarshal([]byte(string(str)), &outPutData); err == nil {
		fmt.Println(outPutData)
	} else {
		fmt.Println(err)
		panic(err)
	}
	return outPutData
}

type EntranceParam struct {
	RequestUrl     string `json:"request_url" description:"请求地址"`
	RequestDataUrl string `json:"request_data_url" description:"请求数据"`
	Type           string `json:"type" description:"请求方式"`
	Addr           string `json:"addr" description:"测试平台"`
}

func Entrance(path string) []EntranceParam {
	entrance := common.ReadJson(path)
	requestData := make([]EntranceParam, 0)
	err := json.Unmarshal([]byte(entrance), &requestData)
	common.CheckError(err)
	return requestData
}
