package parser

import (
	bytes2 "bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/httplib"
	"time"
)

type ParserHomeUrl struct {

}

func (p *ParserHomeUrl) ParserItem(urlList1 chan string, urlList2 chan string) {
ck:
	for {
		select {
		case url, ok := <-urlList1:
			if ok {


			// 爬虫主体Start
				a := 0
			bk:
				bytes, e := httplib.EuUserGet(url)
				if e != nil {
					clog.Println("# 第二级: 任务错误，正在尝试第" + string(a) + "次")
					clog.Println(e.Error())
					a += 1
					time.Sleep(time.Second * 10)
					goto bk
				}

				// 没有异常进行 解析url
				document, e := goquery.NewDocumentFromReader(bytes2.NewReader(bytes))
				if e != nil {
					panic(e.Error())
				}

				document.Find("ul[class='products']").Find("li").Each(func(i int, selection *goquery.Selection) {
					val, exists := selection.Find("a").Attr("href")
					if exists {
						urlList2 <- val
					}
				})

			// end

			} else {
				close(urlList2)
				break ck
			}
		}
	}
}
