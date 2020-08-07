# JccApiTest
* 1.entrance.json为配置入口文件
* 2.json目录为要测试的接口请求参数和响应结果,与入口文件关联
* 3.运行tpl.go,在check目录生成测试代码
* 4.执行go test -v gocheck_test.go开始执行测试用例