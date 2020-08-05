package common

type Output struct {
	Code   string      `json:"code" description:"编码"`
	Msg    string      `json:"msg" description:"提示信息"`
	Data   interface{} `json:"data" description:"数据"`
	Custom interface{} `json:"custom" description:"自定义信息"`
	Token  string      `json:"token" description:"Token"`
}
