package check_test

import (
	"JccApiTest/go_check/common"
	"JccApiTest/go_check/handle"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	. "gopkg.in/check.v1"
	"strconv"
	"testing"
)

var PmToken string = "c0ee9697-9242-4dd2-a0e9-16d20b3b4a59"

var a int = 1

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) SetUpSuite(c *C) {
	//s.TestHttpPost(c)
	str3 := "第1次套件开始执行"
	fmt.Println(str3)
	//c.Skip("Skip TestSuite")
}

func (s *MySuite) TearDownSuite(c *C) {
	str4 := "第一次套件执行完毕"
	fmt.Println(str4)
}

func (s *MySuite) SetUpTest(c *C) {
	str1 := "第" + strconv.Itoa(a) + "条用例开始执行"
	fmt.Println(str1)
}

func (s *MySuite) TearDownTest(c *C) {
	str2 := "第" + strconv.Itoa(a) + "条用例执行完成"
	fmt.Println(str2)
	a = a + 1
}

func (s *MySuite) TestHttpRequest(c *C) {
	entrance := handle.Entrance("../entrance.json")
	if len(entrance) == 0 {
		return
	}
	for _, env := range entrance {
		readData := common.ReadJson("../json/" + env.RequestDataUrl)
		requestData := make([]map[string]string, 0)
		err := json.Unmarshal([]byte(readData), &requestData)
		common.CheckError(err)
		for _, v := range requestData {
			url := fmt.Sprintf("%v/%s", common.PmTestUrl, env.RequestUrl)
			var req *httplib.BeegoHTTPRequest
			switch env.Type {
			case "POST":
				req = httplib.Post(url)
			case "GET":
				req = httplib.Get(url)
			default:
				panic("数据错误")
			}
			switch env.Addr {
			case common.PmAddr:
				req.Header("PmToken", PmToken)
			case common.CsAddr:
				req.Header("PmToken", PmToken)
			}
			for si, sv := range v {
				req.Param(si, sv)
			}
			outPutData := handle.HandleReq(req)
			var code = outPutData.Code
			PmToken = outPutData.Token
			c.Assert(code, Equals, common.SuccessCode) //模拟成功的断言
		}
	}
}

//func (s *MySuite) TestHttpGet(c *C) {
//	getUrl := fmt.Sprintf("%v/pm_member/select_members", common.PmTestUrl)
//	readData := common.ReadJson("../json/select_member.json")
//	requestData := make([]map[string]string, 0)
//	err := json.Unmarshal([]byte(readData), &requestData)
//	common.CheckError(err)
//	for _, v := range requestData {
//		req := httplib.Get(getUrl)
//		for si, sv := range v {
//			req.Param(si, sv)
//		}
//		req.Header("PmToken", PmToken)
//		outPutData := handle.HandleReq(req)
//		var code = outPutData.Code
//		c.Assert(code, Equals, common.SuccessCode) //模拟失败的断言
//	}
//}
