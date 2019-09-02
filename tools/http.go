package tools

import (
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/httplib"
	"net/http"
	"net/url"
	"time"
)


func HttpGet(urls string) ([]byte, error) {

	resp, e := httplib.Get(urls).SetUserAgent(easyutils.ReptileGetUserAgent()).SetProxy(func(request *http.Request) (urls *url.URL, e error) {
		u := new(url.URL)
		u.Scheme = "http"
		u.Host = "127.0.0.1:8001" //蓝灯的http代理地址
		return u, nil
	}).Bytes()

	time.Sleep(time.Second * 3)

	if e != nil {
		return nil,e
	}

	return resp,nil
}
