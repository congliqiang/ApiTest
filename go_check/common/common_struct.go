package common

type Output struct {
	Code   string      `json:"code" description:"编码"`
	Msg    string      `json:"msg" description:"提示信息"`
	Data   interface{} `json:"data" description:"数据"`
	Custom interface{} `json:"custom" description:"自定义信息"`
	Token  string      `json:"token" description:"Token"`
}

type AgentOutput struct {
	Code   int         `json:"code" description:"编码"`
	Msg    string      `json:"msg" description:"提示信息"`
	Data   interface{} `json:"data" description:"数据"`
	Custom interface{} `json:"custom" description:"自定义信息"`
}

type RequestJsonStruct struct {
	RequestUrl  string            `json:"request_url"`
	RequestData map[string]string `json:"request_data"`
}
