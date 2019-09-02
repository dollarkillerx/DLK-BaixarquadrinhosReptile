package test

import (
	"DLK-BaixarquadrinhosReptile/tools"
	bytes2 "bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/httplib"
	"io/ioutil"
	"log"
	"strings"
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	gg("http://baixarquadrinhos.com/Hq-Quadrinho/ler-manga-online-nanatsu-no-taizai-284-ou-baixar-em-cbr-e-pdf/")
	anshtml()
}

func gg(url string) {
	a := 0
bk:
	s, e := tools.AnalysisHtml(url)
	if e != nil {
		clog.Println("# 第三级: 任务错误，正在尝试第" + string(a) + "次")
		clog.Println(e.Error())
		a += 1
		time.Sleep(time.Second * 10)
		goto bk
	}

	//e = ioutil.WriteFile("info.html", []byte(s), 00666)
	//if e != nil {
	//	panic(e.Error())
	//}
	//
	//document, e := goquery.NewDocumentFromReader(strings.NewReader(s))
	//
	//if e != nil {
	//	panic(e.Error())
	//}
	//
	//val, exists := document.Find("img").Attr("src")
	//if exists {
	//	log.Println(val)
	//} else {
	//	log.Println("no")
	//}

	document, e := goquery.NewDocumentFromReader(strings.NewReader(s))
	if e != nil {
		panic(e.Error())
	}

	val, exists := document.Find("div[class='images']").Find("img[data-pagespeed-loaded='1']").Attr("src")
	if exists {
		log.Println(val)
	}

	text := document.Find("div[class='summary entry-summary']").Find("h6").Text()
	log.Println(text)

	s, b := document.Find("div[class='links-download']").Find("a").Attr("href")
	if b {
		log.Println(s)
	}

	get, e := httplib.EuUserGet(s)
	if e != nil {
		panic(e.Error())
	}

	e = ioutil.WriteFile("one.zip", get, 00666)
	if e != nil {
		panic(e.Error())
	}

}

func anshtml() {
	bytes, e := ioutil.ReadFile("info.html")
	if e != nil {
		panic(e.Error())
	}
	document, e := goquery.NewDocumentFromReader(bytes2.NewBuffer(bytes))
	if e != nil {
		panic(e.Error())
	}

	val, exists := document.Find("div[class='images']").Find("img[data-pagespeed-loaded='1']").Attr("src")
	if exists {
		log.Println(val)
	}

	text := document.Find("div[class='summary entry-summary']").Find("h6").Text()
	log.Println(text)

	s, b := document.Find("div[class='links-download']").Find("a").Attr("href")
	if b {
		log.Println(s)
	}

	get, e := httplib.EuUserGet(s)
	if e != nil {
		panic(e.Error())
	}

	e = ioutil.WriteFile("one.zip", get, 00666)
	if e != nil {
		panic(e.Error())
	}
}
