package test

import (
	"DLK-BaixarquadrinhosReptile/defs"
	"DLK-BaixarquadrinhosReptile/tools"
	bytes2 "bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/httplib"
	"io/ioutil"
	"log"
	"strings"
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	book := gg("http://baixarquadrinhos.com/Hq-Quadrinho/ler-manga-online-nanatsu-no-taizai-284-ou-baixar-em-cbr-e-pdf/")

	log.Println(book)
	//anshtml()

	//kk()

	//test()
}

func gg(url string) *defs.Book {
	a := 0
bk:
	s, e := tools.AnalysisHtml(url)
	//b1, e := httplib.EuUserGet(url)

	if e != nil {
		clog.Println("# 第三级: 任务错误，正在尝试第" + string(a) + "次")
		clog.Println(e.Error())
		a += 1
		time.Sleep(time.Second * 10)
		goto bk
	}

	document, e := goquery.NewDocumentFromReader(strings.NewReader(s))
	//document, e := goquery.NewDocumentFromReader(strings.NewReader(string(b1)))
	if e != nil {
		panic(e.Error())
	}

	val, exists := document.Find("div[class='images']").Find("img[data-pagespeed-loaded='1']").Attr("src")
	log.Println(val)

	text := document.Find("div[class='summary entry-summary']").Find("h6").Text()
	log.Println(text)

	s, b := document.Find("div[class='links-download']").Find("a").Attr("href")
	if b {
		log.Println(s)
	}

	a = 0
bk1:
	by, e := httplib.EuUserGet(s)
	if e != nil {
		clog.Println("# 第三级: 任务错误（下载zip），正在尝试第" + string(a) + "次")
		clog.Println(e.Error())
		a += 1
		time.Sleep(time.Second * 10)
		goto bk1
	}

bk2:
	by2, e := httplib.EuUserGet(val)
	if e != nil {
		if exists {
			clog.Println("# 第三级: 任务错误（下载img），正在尝试第" + string(a)  + "次")
			clog.Println(e.Error())
			a += 1
			time.Sleep(time.Second * 10)
			goto bk2
		}
	}

	filepath := "./" + strings.TrimSpace(text) + "/" + easyutils.SuperRand()

	e = easyutils.DirPing(filepath)
	if e != nil {
		panic(e.Error())
	}

	e = ioutil.WriteFile(filepath+"/one.zip", by, 00666)
	if e != nil {
		panic(e.Error())
	}

	postfix, e := easyutils.FileGetPostfix(val)
	if e != nil {
		panic(e.Error())
	}
	clog.Println(val)

	imgpath := filepath + "/" + easyutils.SuperRand() + "." + postfix
	e = ioutil.WriteFile(imgpath, by2, 00666)
	if e != nil {
		panic(e.Error())
	}

	data := defs.Book{
		Img:  imgpath,
		Path: filepath + "/one.zip",
		Name: strings.TrimSpace(text),
	}

	return &data

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

func kk() {
	e := easyutils.DirPing("./hello/pc")
	if e != nil {
		panic(e.Error())
	}
}

func test() {
	url := "http://baixarquadrinhos.com/Hq-Quadrinho/ler-manga-online-nanatsu-no-taizai-284-ou-baixar-em-cbr-e-pdf/"
	bytes, e := httplib.EuUserGet(url)
	if e != nil {
		panic(e.Error())
	}

	ioutil.WriteFile("hellp.html", bytes, 00666)
}
