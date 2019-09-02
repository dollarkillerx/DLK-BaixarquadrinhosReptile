package parser

import (
	"DLK-BaixarquadrinhosReptile/defs"
	"DLK-BaixarquadrinhosReptile/tools"
	"github.com/PuerkitoBio/goquery"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

// 解析主体
type ParserInfo struct {
}

// 主要: 第一解析 ，第二 获取首图，下载地址，名称 入库
func (p *ParserInfo) ParserItem(data chan string, data2 chan interface{}) {
	// 获取首图 并下载
	for {
		select {
		case url, ok := <-data:
			if ok {
				log.Print("收到任务正在执行 下载")
				gg(url)
			} else {
				data2 <- 1
			}
		}
	}
}

func gg(url string) *defs.Book {
	a := 0
bk:
	//s, e := tools.AnalysisHtml(url)

	b1, e := tools.HttpGet(url)

	if e != nil {
		clog.Println("# 第三级: 任务错误，正在尝试第" + string(a) + "次")
		clog.Println(e.Error())
		a += 1
		time.Sleep(time.Second * 10)
		goto bk
	}

	//document, e := goquery.NewDocumentFromReader(strings.NewReader(s))
	document, e := goquery.NewDocumentFromReader(strings.NewReader(string(b1)))
	if e != nil {
		panic(e.Error())
	}



	text := document.Find("div[class='summary entry-summary']").Find("h6").Text()
	log.Println(text)

	s, b := document.Find("div[class='links-download']").Find("a").Attr("href")
	if b {
		log.Println(s)
	}

	a = 0
bk1:

	by, e := tools.HttpGet(s)
	if e != nil {
		if a > 10 {
			return nil
		}
		clog.Println("# 第三级: 任务错误（下载zip），正在尝试第" + string(a) + "次")
		clog.Println(e.Error())
		a += 1
		time.Sleep(time.Second * 10)
		goto bk1
	}

//bk2:
//	by2, e := tools.HttpGet(s)
//	if e != nil {
//		if exists {
//			clog.Println("# 第三级: 任务错误（下载img），正在尝试第" + string(a) + "次")
//			clog.Println(e.Error())
//			a += 1
//			time.Sleep(time.Second * 10)
//			goto bk2
//		}
//	}

	filepath := "./img/" + strings.TrimSpace(text) + "/" + easyutils.SuperRand()

	e = easyutils.DirPing(filepath)
	if e != nil {
		panic(e.Error())
	}

	e = ioutil.WriteFile(filepath+"/one.zip", by, 00666)
	if e != nil {
		panic(e.Error())
	}

	//postfix, e := easyutils.FileGetPostfix(val)
	//if e != nil {
	//	panic(e.Error())
	//}
	//clog.Println(val)
	//
	//imgpath := filepath + "/" + easyutils.SuperRand() + "." + postfix
	//e = ioutil.WriteFile(imgpath, by2, 00666)
	//if e != nil {
	//	panic(e.Error())
	//}

	data := defs.Book{
		//Img:  imgpath,
		Path: filepath + "/one.zip",
		Name: strings.TrimSpace(text),
	}

	return &data
}
