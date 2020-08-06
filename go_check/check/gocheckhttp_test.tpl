package check_test

import (
	"JccApiTest/go_check/handle"
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

{{.NewFunc}}
