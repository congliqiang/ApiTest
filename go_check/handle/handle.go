package handle

import (
	"JccApiTest/go_check/common"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
)

func HandleReq(req *httplib.BeegoHTTPRequest) *common.Output {
	str, err := req.String()
	if err != nil {
		panic(err)
	}

	outPutData := new(common.Output)
	if err := json.Unmarshal([]byte(string(str)), &outPutData); err == nil {
		fmt.Println(outPutData)
	} else {
		fmt.Println(err)
		panic(err)
	}
	return outPutData
}
