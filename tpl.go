package main

import (
	"JccApiTest/go_check/common"
	"JccApiTest/go_check/handle"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"os"
	"strconv"
	"text/template"
	"time"
)

func main() {
	tpl, err := template.ParseFiles("go_check/check/gocheckhttp_test.tpl")
	if err != nil {
		panic(err)
	}
	br := new(bytes.Buffer)

	data := make(map[string]interface{})
	entrance := handle.Entrance("./go_check/entrance.json")
	if len(entrance) == 0 {
		return
	}
	dataFunc := ""
	for eni, env := range entrance {
		readData := common.ReadJson("./go_check/json/" + env.RequestDataUrl)
		requestData := make([]map[string]string, 0)
		err := json.Unmarshal([]byte(readData), &requestData)
		common.CheckError(err)
		fmt.Println(requestData)
		for i, v := range requestData {
			url := fmt.Sprintf("%v/%s", common.PmTestUrl, env.RequestUrl)
			var req *httplib.BeegoHTTPRequest
			switch env.Type {
			case "Post":
				req = httplib.Post(url)
			case "Get":
				req = httplib.Get(url)
			default:
				panic("数据错误")
			}
			switch env.Addr {
			case common.PmAddr:
				req.Header("PmToken", common.PmToken)
			case common.CsAddr:
				req.Header("Token", common.PmToken)
			}
			var str = ""
			for si, sv := range v {
				if sv != "" {
					str += `req.Param("` + si + `","` + sv + `")
			`
				}

			}
			var funcName = "TestHttpRequest"
			funcName += strconv.Itoa(eni)
			funcName += strconv.Itoa(i)
			funcName += strconv.Itoa(time.Now().Nanosecond())
			dataFunc += fmt.Sprintf(`func (s *MySuite) %s(c *C) {
			req := httplib.%s("%s")
			req.Header("%s", "%s")
			%s
			outPutData := handle.HandleReq(req)
			var code = outPutData.Code
			PmToken = outPutData.Token
			c.Assert(code, Equals, common.SuccessCode)
}
`, funcName, env.Type, url, env.Addr, common.PmToken, str)
		}
	}
	data["NewFunc"] = dataFunc
	err = tpl.Execute(br, data)
	if err != nil {
		panic(err)
	}
	f, e := os.Create("./go_check/check/gocheck_test.go")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	_, _ = f.Write([]byte(br.String()))
	//fmt.Println(br.String())
}
