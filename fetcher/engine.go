package fetcher

import "DLK-BaixarquadrinhosReptile/parser"

type ParserUrl interface {
	ParserUrl(chan string)
}

type ParserUrlItem interface {
	ParserItem(chan string, chan string)
}

type ParserItem interface {
	ParserItem(chan string, chan interface{})
}

type ParserEnd interface {
	ParserEnd(chan interface{})
}

// 爬虫主体
type Reptile struct {
	churl chan string      // url 主页url任务
	chin1 chan interface{} // 第二级
	chin2 chan interface{} // 第三级

	parserUrl ParserUrl  // url 生成器
	parser1   ParserUrlItem // 第二级 处理
	parser2   ParserItem // 第三级处理
}

// 中央控制
func ReptileEngine() {
	generateUrl := &parser.GenerateUrl{}
	parser1 := &parser.ParserHomeUrl{}

	data := Reptile{
		churl:     make(chan string, 1000),
		chin1:     make(chan interface{}, 1000),
		chin2:     make(chan interface{}, 1000),
		parserUrl: generateUrl,
		parser1:   parser1,
		parser2:   nil,
	}



}
