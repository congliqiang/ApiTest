package main

import (
	"JccApiTest/go_check/common"
	"JccApiTest/go_check/handle"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"math/rand"
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
			if env.Addr == common.AgentAddr {
				url = fmt.Sprintf("%v/%s", common.AgentUrl, env.RequestUrl)
			}
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
			case common.AgentAddr:
				req.Header("jcc-path", common.Jcc_Path)
			}
			var str = ""
			var task_id = ""
			var checkErrorMsg = ""
			for si, sv := range v {
				if si == "response" {
					checkErrorMsg = sv
					continue
				}
				// 随机变量 适用于增加场景
				if sv == "[randName]" {
					sv = randSeq() // 实际生成 3位随机数+当前10位时间戳的字符串
				}
				// 禅道号遍历
				if si == "task_id" {
					task_id = sv
					continue
				}
				if sv != "" {
					str += "req.Param(`" + si + "`,`" + sv + "`" + `)
			`
				} else {
					str += `req.Param("` + si + `","")
			`
				}

			}
			var funcName = "TestHttpRequest"
			funcName += strconv.Itoa(eni)
			funcName += strconv.Itoa(i)
			funcName += strconv.Itoa(time.Now().Nanosecond())
			dataFunc += fmt.Sprintf(`func (s *MySuite) %s(c *C) {
			fmt.Println("开始执行 禅道测试用例编号为 【%s】的任务")
			req := httplib.%s("%s")
			req.Header("%s", "%s")
			%s
			outPutData := handle.HandleReq(req)
			var msg = outPutData["msg"]
			c.Assert(msg, Equals, "%s")
}
`, funcName, task_id, env.Type, url, env.Addr, common.PmToken, str, checkErrorMsg)
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
func randSeq() (b string) {
	p := []int{3, 5, 6, 7, 8, 9}
	var d string
	a := time.Now().UnixNano()
	rand.Seed(a) //(88-15 )+15
	for i := 0; i < 10; i++ {
		d = d + strconv.Itoa(rand.Intn(10))
	}
	l := rand.Intn(5)
	o := strconv.Itoa(p[l])
	return "1" + o + d
}
