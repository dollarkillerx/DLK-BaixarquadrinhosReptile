package test

import (
	bytes2 "bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/httplib"
	"log"
	"testing"
)

func TestParserHome(t *testing.T) {
	jx("http://baixarquadrinhos.com/page/2/")
}

func jx(url string) string {
	a := 0
bk:
	bytes, e := httplib.EuUserGet(url)
	if e != nil {
		clog.Println("# 第二级: 任务错误，正在尝试第" + string(a) + "次")
		clog.Println(e.Error())
		a += 1
		goto bk
	}

	document, e := goquery.NewDocumentFromReader(bytes2.NewReader(bytes))
	if e != nil {
		panic(e.Error())
	}

	document.Find("ul[class='products']").Find("li").Each(func(i int, selection *goquery.Selection) {
		val, exists := selection.Find("a").Attr("href")
		if exists {
			log.Println(val)
		}
	})
	return ""
}
