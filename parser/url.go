package parser

import (
	"strconv"
)

type GenerateUrl struct {
}

// url 生成器
func (p *GenerateUrl) ParserUrl(chUrl chan string) {
	for i := 2; i <= 83; i++ {
		url := "http://baixarquadrinhos.com/page/" + strconv.Itoa(i) + "/"
		chUrl <- url
	}
	close(chUrl)
}
